package address

import (
	"SentiService/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupAddressRoutes(api *mux.Router, handler *Handler) {
	address := api.PathPrefix("/addresses").Subrouter()

	protected := address.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("", handler.Create).Methods("POST")
	protected.HandleFunc("", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
