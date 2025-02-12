package srv

import (
	"chetam/cfg"
	"chetam/internal/service"
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/riandyrn/otelchi"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	otelchimetric "github.com/riandyrn/otelchi/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var tracer oteltrace.Tracer

const (
	serverName = "chetam"
)

type Config struct {
	cfg.SRV `envPrefix:"SRV_"`
}

type Server struct {
	cfg     Config
	lg      *slog.Logger
	service *service.Service
}

func New(logger *slog.Logger, service *service.Service) Server {
	c := Config{}
	if err := cfg.Parse(&c); err != nil {
		panic(fmt.Errorf("no env for server: %w", err))
	}
	return Server{
		cfg:     c,
		lg:      logger,
		service: service,
	}
}

func (s *Server) Start() {
	tp := initTracerProvider()
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	tracer = otel.Tracer("mux-server")

	mp := initMeter()
	otel.SetMeterProvider(mp)

	baseCfg := otelchimetric.NewBaseConfig(serverName, otelchimetric.WithMeterProvider(mp))
	r := chi.NewRouter()

	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(
		otelchi.Middleware(serverName, otelchi.WithChiRoutes(r)),
		otelchimetric.NewRequestDurationMillis(baseCfg),
		otelchimetric.NewRequestInFlight(baseCfg),
		otelchimetric.NewResponseSizeBytes(baseCfg),
	)
	r.Route("/users", func(r chi.Router) {
		r.Get("/", s.getUsersHandler)
		r.Post("/add", s.addUserHandler)
		r.Patch("/{user}", s.updateUserHandler)
		r.Delete("/{user}", s.deleteUserHandler)
	})

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/{user}", s.getTasksHandler)
		r.Post("/start", s.addTaskHandler)
		r.Post("/end", s.endTaskHandler)
	})

	//r.Get("/swagger/*", httpSwagger.Handler(
	//	httpSwagger.URL(fmt.Sprintf("http://%s/swagger/doc.json", s.cfg.Address)),
	//))

	s.lg.Info("starting server",
		slog.String("address", ":"+s.cfg.PORT))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:         ":" + s.cfg.PORT,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			s.lg.Debug("server error",
				slog.String("error", err.Error()),
			)

			if !errors.Is(err, http.ErrServerClosed) {
				s.lg.Error("failed to start server")
			}
		}
	}()

	s.lg.Info("server started")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	<-done
	s.lg.Info("shutting down server")

	if err := srv.Shutdown(ctx); err != nil {
		s.lg.Error("failed to shutdown server")
	}

	s.lg.Info("server shutdown")
}

func initTracerProvider() *sdktrace.TracerProvider {
	exporter, err := stdout.New(stdout.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}
	res, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String("mux-server"),
		),
	)
	if err != nil {
		log.Fatalf("unable to initialize resource due: %v", err)
	}
	return sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)
}

func initMeter() *sdkmetric.MeterProvider {
	exp, err := stdoutmetric.New()
	if err != nil {
		log.Fatal(err)
	}

	return sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(exp)),
	)
}
