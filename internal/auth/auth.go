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
	Login string `json:"login"`
	jwt.RegisteredClaims
}

func (a *Auth) generateJWT(login string) (string, error) {
	claims := Claims{
		Login: login,
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

//
//func (a *Auth) GenerateCode(w http.ResponseWriter, r *http.Request) {
//	baseUrl := "https://sms.ru/sms/send"
//
//	u, _ := url.Parse(baseUrl)
//	params := url.Values{}
//	params.Add("api_id", a.cfg.Jwt.Sms)
//	params.Add("to", a.cfg.Jwt.Phone)
//	params.Add("msg", "hi")
//	params.Add("json", "1")
//	u.RawQuery = params.Encode()
//	// Отправляем GET-запрос
//	resp, err := http.Get(u.String())
//	if err != nil {
//		log.Fatal("Error sending GET request:", err)
//	}
//	defer resp.Body.Close()
//	// Чтение ответа
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatal("Error reading response body:", err)
//	}
//	fmt.Println("Response:", string(body))
//}
