package auth

import (
	"chetam/internal/config"
	"chetam/internal/transport/repository"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Service struct {
	cfg  *config.Config
	lg   *slog.Logger
	repo *repository.Repository
}

func New(cfg *config.Config, lg *slog.Logger, repo *repository.Repository) *Service {
	service := &Service{
		cfg:  cfg,
		lg:   lg,
		repo: repo,
	}

	return service
}

type Claims struct {
	Login string `json:"login"`
	jwt.RegisteredClaims
}

func (s *Service) GenerateCode(w http.ResponseWriter, r *http.Request) {
	baseUrl := "https://sms.ru/sms/send"

	u, _ := url.Parse(baseUrl)
	params := url.Values{}
	params.Add("api_id", s.cfg.Jwt.Sms)
	params.Add("to", s.cfg.Jwt.Phone)
	params.Add("msg", "hi")
	params.Add("json", "1")
	u.RawQuery = params.Encode()

	// Отправляем GET-запрос
	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal("Error sending GET request:", err)
	}
	defer resp.Body.Close()

	// Чтение ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}
	fmt.Println("Response:", string(body))

}

func (s *Service) Register(email, login, password string) (string, error) {
	user, err := s.repo.CreateUser(email, login, password)
	if err != nil {
		return "", err
	}

	token, err := s.generateJWT(user.Login)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) Login(login, password string) (string, error) {
	user, err := s.repo.FindUserByLogin(login)
	if err != nil {
		s.lg.Warn("user not found",
			slog.String("error", err.Error()))

		return "", err
	} else if user.Password != password {
		return "", fmt.Errorf("password incorrect")
	}

	token, err := s.generateJWT(user.Login)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) generateJWT(login string) (string, error) {
	claims := Claims{
		Login: login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.cfg.Jwt.SecretKey))
	if err != nil {
		return err.Error(), err
	}

	return tokenString, nil
}

func (s *Service) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.Jwt.SecretKey), nil
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
