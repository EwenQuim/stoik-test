package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

const (
	timeout = 10 * time.Second // Timeout for requests
)

type Server struct {
	http.Server
	Router *chi.Mux // Must be moved to http.Server when serving
}

func main() {
	r := chi.NewMux()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	s := &Server{
		Server: http.Server{
			Addr:         ":8090",
			ReadTimeout:  timeout,
			WriteTimeout: timeout,
			Handler:      r,
		},
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
