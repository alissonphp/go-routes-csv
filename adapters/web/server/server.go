package server

import (
	"github.com/gorilla/mux"
	"go-best-route/adapters/web/handler"
	"go-best-route/application"
	"log"
	"net/http"
	"time"
)

type Webserver struct {
	Service application.RouteServiceInterface
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Server() {
	r := mux.NewRouter()
	handler.SetRouteHandlers(r, w.Service)
	http.Handle("/", r)
	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
