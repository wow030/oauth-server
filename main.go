package main

import (
	"log"
	"net/http"
	"oauth-server/service"
)

func main() {
	server := service.CreateService()

	err := http.ListenAndServeTLS(":8081", "server.crt", "server.key", server.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
