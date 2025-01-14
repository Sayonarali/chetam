package auth

import (
	"chetam/cfg"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

type Config struct {
	cfg.JWT `envPrefix:"JWT_"`
}

type Service struct {
	cfg Config
}

func NewAuthService(c Config) Service {
	return Service{
		cfg: c,
	}
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s Service) Register(w http.ResponseWriter, r *http.Request) {
	//var req struct {
	//	Username string `json:"username"`
	//	Password string `json:"password"`
	//}
	//
	//if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	//	http.Error(w, "Invalid request", http.StatusBadRequest)
	//	return
	//}
}

func (s Service) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	token, err := s.generateJWT(req.Username)
	if err != nil {
		http.Error(w, "Invalid request generate JWT", http.StatusBadRequest)
	}

	resp := struct {
		Token string `json:"token"`
	}{Token: token}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s Service) generateJWT(username string) (string, error) {
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.cfg.SecretKey))
	if err != nil {
		return err.Error(), err
	}

	return tokenString, nil
}

func (s Service) validateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return s.cfg.SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
