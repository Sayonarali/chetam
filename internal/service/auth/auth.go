package auth

import (
	"chetam/cfg"
	"chetam/internal/service/repository"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"time"
)

type Config struct {
	cfg.JWT `envPrefix:"JWT_"`
}

type Service struct {
	lg               *slog.Logger
	cfg              Config
	repositoryKeeper repository.Keeper
}

func NewAuthService(c Config, rk repository.Keeper, lg *slog.Logger) Service {
	service := Service{
		lg:               lg,
		cfg:              c,
		repositoryKeeper: rk,
	}
	return service
}

type Claims struct {
	Login string `json:"login"`
	jwt.RegisteredClaims
}

func (s Service) GenerateCode(w http.ResponseWriter, r *http.Request) {
	baseUrl := "https://sms.ru/sms/send"

	u, _ := url.Parse(baseUrl)
	params := url.Values{}
	params.Add("api_id", s.cfg.Sms)
	params.Add("to", s.cfg.Phone)
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

func (s Service) Register(email, login, password string) (string, error) {
	user, err := s.repositoryKeeper.CreateUser(email, login, password)
	if err != nil {
		return "", err
	}
	token, err := s.generateJWT(user.Login)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s Service) Login(login, password string) (string, error) {
	user, err := s.repositoryKeeper.FindUserByLogin(login)
	if err != nil {
		s.lg.Warn(err.Error())
		return "", err
	} else if user.Password != password {
		return "", errors.New("wrong password")
	}

	token, err := s.generateJWT(user.Login)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s Service) generateJWT(login string) (string, error) {
	claims := Claims{
		Login: login,
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

func (s Service) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.SecretKey), nil
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
