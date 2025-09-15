package main

import (
	"go-back/internal/http/handler"
	"go-back/internal/http/router"
)

func main() {

	r := router.NewRouter()
	handler.HandleRequests(r)
	r.Run(":1111")
}
