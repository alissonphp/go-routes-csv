package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-best-route/adapters/dto"
	"go-best-route/application"
	"net/http"
)

func SetRouteHandlers(r *mux.Router, service application.RouteServiceInterface) {
	r.Handle("/routes", createRoute(service)).Methods("POST", "OPTIONS")
	r.Handle("/routes/cheapest", bestRoute(service)).Queries("from", "{from}").Queries("to", "{to}").Methods("GET", "OPTIONS")
}

func createRoute(service application.RouteServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var createRouteDto dto.CreateRoute

		err := json.NewDecoder(r.Body).Decode(&createRouteDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		route, err := service.Save(createRouteDto.From, createRouteDto.To, createRouteDto.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(route)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}

func bestRoute(service application.RouteServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		from := r.FormValue("from")
		to := r.FormValue("to")
		best, err := service.SearchBest(from, to)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(best)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
