package main

import (
	"go-back/internal/http/handler"
	"go-back/internal/http/router"
)

func main() {

	r := router.NewRouter()
	handler.HandleRequests(r)
	err := r.Run(":1111")
	if err != nil {
		panic(err)
	}
}
