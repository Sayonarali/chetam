package auth

import (
	"chetam/internal/config"
	"chetam/internal/repository"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log/slog"
	"time"
)

type Auth struct {
	cfg  *config.Config
	lg   *slog.Logger
	repo *repository.Repository
}

func New(cfg *config.Config, lg *slog.Logger, repo *repository.Repository) *Auth {
	return &Auth{
		cfg:  cfg,
		lg:   lg,
		repo: repo,
	}
}

type Claims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

func (a *Auth) generateJWT(id int) (string, error) {
	claims := Claims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.cfg.Jwt.SecretKey))
	if err != nil {
		return err.Error(), err
	}

	return tokenString, nil
}

func (a *Auth) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.cfg.Jwt.SecretKey), nil
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
