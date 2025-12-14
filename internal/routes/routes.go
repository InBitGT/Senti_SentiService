package routes

import (
	"SentiService/internal/middleware"
	"SentiService/internal/modules/address"

	"github.com/gorilla/mux"
)

type RouteHandlers interface {
	GetAddressHandler() *address.Handler
}

func SetupRoutes(router *mux.Router, handlers RouteHandlers) {
	router.Use(middleware.ContentTypeJSON)
	router.Use(middleware.Logger)
	router.Use(middleware.Recovery)

	api := router.PathPrefix("/api").Subrouter()

	address.SetupAddressRoutes(api, handlers.GetAddressHandler())
}
