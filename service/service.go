package service

import (
	"net/http"
)

type Service struct {
	Mux *http.ServeMux
}

func CreateService() *Service {
	service := Service{}

	service.Mux = http.NewServeMux()

	service.Mux.HandleFunc("/token", TokenHandler)

	return &service
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("token"))
}
