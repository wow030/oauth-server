package main

import (
	"log"
	"net/http"
	"oauth-server/service"
)

func main() {
	server := service.CreateService()

	err := http.ListenAndServe(":8001", server.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
