package model

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	Token string `json:"token,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token,omitempty"`
}

type GetRoutesListResponse struct {
	Token string `json:"token,omitempty"`
}
