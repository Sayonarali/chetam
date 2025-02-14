package handlers

import (
	"log/slog"
	"net/http"
)

type AuthInterface interface {
}

func Register(logger *slog.Logger, service AuthInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func Auth(logger *slog.Logger, service AuthInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
