package server

import (
	"chetam/internal/config"
	"chetam/internal/model"
	"chetam/internal/server/handlers"
	"chetam/internal/services"
	"context"
	"errors"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	cfg      *config.Config
	lg       *slog.Logger
	services *services.Services
}

func New(lg *slog.Logger, cfg *config.Config, services *services.Services) *Server {
	return &Server{
		cfg:      cfg,
		lg:       lg,
		services: services,
	}
}

func (s *Server) Run() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.POST("/auth/register", handlers.Register(s.lg, s.services.Auth))
	e.POST("/auth/login", handlers.Login(s.lg, s.services.Auth))

	apiGroup := e.Group("/api/v1")
	//apiGroup.Use(jwtMiddleware(s.cfg))

	pointsGroup := apiGroup.Group("/points")
	pointsGroup.GET("", handlers.GetUser(s.lg))

	routesGroup := apiGroup.Group("/routes")
	routesGroup.GET("", handlers.GetRoutesList(s.lg, s.services.Route))
	routesGroup.GET("/:id", handlers.GetRoute(s.lg, s.services.Route))

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

func jwtMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	config := jwtConfig(cfg)
	return echojwt.WithConfig(config)
}

func jwtConfig(cfg *config.Config) echojwt.Config {
	return echojwt.Config{
		SigningKey: []byte(cfg.Jwt.SecretKey),
		ContextKey: "token",
		ErrorHandler: func(c echo.Context, err error) error {
			e := model.Error{
				Error: "failed to validate token",
			}
			slog.Info("failed to validate token",
				slog.String("error", err.Error()),
			)

			return c.JSON(http.StatusUnauthorized, e)
		},
	}
}
