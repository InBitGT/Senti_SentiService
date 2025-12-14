package server

import "SentiService/internal/modules/address"

func (h *Handlers) GetAddressHandler() *address.Handler {
	return h.Address
}
