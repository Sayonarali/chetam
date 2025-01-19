package main

import (
	"chetam/cmd/factory"
	chetamApiv1 "chetam/pkg/chetamApi/v1"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	router, err := factory.InitializeRouter()

	if err != nil {
		fmt.Println(err.Error())
	}

	r := chi.NewRouter()
	handler := chetamApiv1.HandlerFromMux(router, r)

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(fmt.Errorf("failed to start server: %w", err))
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("failed to load .env: %w", err))
	}
}
