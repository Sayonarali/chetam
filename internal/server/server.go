package server

import (
	"chetam/internal/config"
	"chetam/internal/model"
	chetamApiv1 "chetam/internal/server/client/v1"
	"chetam/internal/server/handlers"
	"context"
	"errors"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	oteltrace "go.opentelemetry.io/otel/trace"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var tracer oteltrace.Tracer

const (
	serverName = "chetam"
)

type Server struct {
	cfg *config.Config
	sh  *handlers.ServerHandler
	lg  *slog.Logger
}

func New(lg *slog.Logger, cfg *config.Config, sh *handlers.ServerHandler) *Server {
	return &Server{
		cfg: cfg,
		sh:  sh,
		lg:  lg,
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

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))
	e.Use(otelecho.Middleware("chetam"))
	//e.Use(jwtMiddleware(s.cfg))

	apiGroup := e.Group("/api", jwtMiddleware(s.cfg))
	apiGroup.Use(jwtMiddleware(s.cfg))
	//chetamApiv1.RegisterHandlers(apiGroup, &chetamApiv1.ServerInterfaceWrapper{
	//	Handler: s.sh,
	//})
	chetamApiv1.RegisterHandlersWithBaseURL(e, s.sh, "")
	//e.GET("/swagger/*", httpSwagger.Handler(
	//	httpSwagger.URL(fmt.Sprintf("http://%s/swagger/doc.json", s.cfg.Address)),
	//))

	s.lg.Info("starting server",
		slog.String("address", ":"+s.cfg.Server.Port))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := e.Start(":" + s.cfg.Server.Port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.lg.Error("failed to start server",
				slog.String("error", err.Error()),
			)
		}
	}()

	s.lg.Info("server started")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	<-done
	s.lg.Info("shutting down server")

	if err := e.Shutdown(ctx); err != nil {
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

func jwtMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	config := jwtConfig(cfg)
	return echojwt.WithConfig(config)
}

func jwtConfig(cfg *config.Config) echojwt.Config {
	e := model.Error{
		Errors: "Ошбика авторизации",
	}

	return echojwt.Config{
		SigningKey: []byte(cfg.Jwt.SecretKey),
		ContextKey: "token",
		ErrorHandler: func(c echo.Context, err error) error {
			slog.Info("failed to validate token",
				slog.String("error", err.Error()),
			)

			return c.JSON(http.StatusUnauthorized, e)
		},
	}
}
