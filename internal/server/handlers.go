package server

import (
	"SentiService/internal/modules/address"

	"gorm.io/gorm"
)

func NewHandlers(db *gorm.DB) *Handlers {
	addrRepo := address.NewRepository(db)
	addrService := address.NewService(addrRepo)
	addrHandler := address.NewHandler(addrService)

	return &Handlers{
		Address: addrHandler,
	}
}
