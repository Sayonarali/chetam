package server

import (
	"chetam/internal/config"
	"chetam/internal/server/handlers"
	"chetam/internal/services"
	"context"
	"errors"
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

type Server struct {
	cfg      *config.Config
	lg       *slog.Logger
	services *services.Services
}

func New(cfg *config.Config, logger *slog.Logger, services *services.Services) *Server {
	return &Server{
		services: services,
		cfg:      cfg,
		lg:       logger,
	}
}

func (s *Server) Run() {
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
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/register", handlers.Register(s.lg, s.services))
		r.Get("/auth", handlers.Auth(s.lg, s.services))
	})

	//r.Get("/swagger/*", httpSwagger.Handler(
	//	httpSwagger.URL(fmt.Sprintf("http://%s/swagger/doc.json", s.cfg.Address)),
	//))

	s.lg.Info("starting server",
		slog.String("address", ":"+s.cfg.Server.Port))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:         ":" + s.cfg.Server.Port,
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
