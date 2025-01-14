package main

import (
	"chetam/cmd/factory"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	chetamService, err := factory.InitializeChetam()

	if err != nil {
		fmt.Println(err.Error())
	}

	chetamService.Execute(r)
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("failed to load .env: %w", err))
	}
}
