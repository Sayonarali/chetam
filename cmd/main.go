package main

import (
	"chetam/cmd/factory"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	chetamService, err := factory.InitializeChetam()

	if err != nil {
		fmt.Println(err.Error())
	}

	chetamService.Execute()
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err.Error())
	}
}
