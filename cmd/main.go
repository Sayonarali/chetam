package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "go1")
//}

var jwtSecret = []byte("your-secret-key")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func generateJWT(username string) (string, error) {
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
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

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	token, err := generateJWT(req.Username)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	resp := struct {
		Token string `json:"token"`
	}{Token: token}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	//connection := fmt.Sprintf(
	//	"postgres://%s:%s@%s:%s/%s?sslmode=disable",
	//	os.Getenv("DB_USER"),
	//	os.Getenv("DB_PASSWORD"),
	//	os.Getenv("DB_HOST"),
	//	os.Getenv("DB_PORT"),
	//	os.Getenv("DB_NAME"),
	//)
	//db, err := sql.Open("postgres", connection)
	//if err != nil {
	//	log.Fatal("Connection db error", err)
	//}
	//defer db.Close()
	//fmt.Println("Connected to database")
	//
	//http.HandleFunc("/", handler)
	//
	//err = http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	fmt.Println("Err", err)
	//}
	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
