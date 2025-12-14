package address

import "github.com/gorilla/mux"

func SetupAddressRoutes(api *mux.Router, handler *Handler) {
	addrRouter := api.PathPrefix("/addresses").Subrouter()

	addrRouter.HandleFunc("", handler.Create).Methods("POST")
	addrRouter.HandleFunc("", handler.GetAll).Methods("GET")
	addrRouter.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	addrRouter.HandleFunc("/{id}", handler.Update).Methods("PUT")
	addrRouter.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
