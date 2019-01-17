package handler

import (
	"github.com/mochisuna/ssh-mysql-sample/domain/service"
)

func New(services *Services) *Handler {
	return &Handler{
		Services: *services,
	}
}

type Services struct {
	StoreService service.StoreService
}

type Handler struct {
	Services
}
