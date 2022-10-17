package service

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"

	"github.com/alexm24/ticket-booking/internal/models"
	"github.com/alexm24/ticket-booking/internal/repository"
)

type IRoutes interface {
	CreateRoute(item models.PostRoute) (models.Route, error)
}

type IPassengers interface {
	CreatePassenger(routeId openapi_types.UUID, item models.PostPassenger) (models.Passenger, error)
}

type Service struct {
	IRoutes
	IPassengers
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		IRoutes:     NewRoutesService(repos),
		IPassengers: NewPassengersService(repos),
	}
}
