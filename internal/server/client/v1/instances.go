package chetamApiv1

// LoginRequest defines model for LoginRequest.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse defines model for LoginResponse.
type LoginResponse struct {
	Token *string `json:"token,omitempty"`
}

// RegisterRequest defines model for RegisterRequest.
type RegisterRequest struct {
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

// RegisterResponse defines model for RegisterResponse.
type RegisterResponse struct {
	Token *string `json:"token,omitempty"`
}

// User defines model for User.
type User struct {
	Email  *string `json:"email,omitempty"`
	Login  *string `json:"login,omitempty"`
	UserId *int    `json:"userId,omitempty"`
}
