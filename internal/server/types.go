package server

import (
	"SentiService/internal/modules/address"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	Router *mux.Router
	DB     *gorm.DB
}

type Handlers struct {
	Address *address.Handler
}
