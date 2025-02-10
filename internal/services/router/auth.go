package router

import (
	chetamApiv1 "chetam/pkg/chetamApi/v1"
	"encoding/json"
	"net/http"
)

func (rt Router) PostApiV1AuthRegister(w http.ResponseWriter, r *http.Request) {
	var req chetamApiv1.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := rt.authService.Register(req.Email, req.Login, req.Password)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	resp := chetamApiv1.RegisterResponse{Token: &token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

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
