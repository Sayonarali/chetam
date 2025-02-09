package main

import (
	"chetam/cmd/factory"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("failed to load .env: %w", err))
	}

	_, err := factory.InitializeRouter()
	if err != nil {
		fmt.Println(err.Error())
	}
}
