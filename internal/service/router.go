package router

import (
	"chetam/internal/service/auth"
	"chetam/internal/service/user"
	chetamApiv1 "chetam/pkg/chetamApi/v1"
	"encoding/json"
	"log/slog"
	"net/http"
)

type Router struct {
	lg          *slog.Logger
	userService *user.Service
	authService *auth.Service
}

func New(
	lg *slog.Logger,
	us *user.Service,
	as *auth.Service,
) Router {
	if as == nil || lg == nil {
		panic("nil argument in new")
	}
	return Router{
		lg:          lg,
		userService: us,
		authService: as,
	}
}

func (rt Router) PostApiV1AuthLogin(w http.ResponseWriter, r *http.Request) {
	var req chetamApiv1.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := rt.authService.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	resp := chetamApiv1.LoginResponse{Token: &token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (rt Router) GetApiV1User(w http.ResponseWriter, r *http.Request) {

}
