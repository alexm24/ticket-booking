package service

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"

	"github.com/alexm24/ticket-booking/internal/models"
	"github.com/alexm24/ticket-booking/internal/repository"
)

type PassengersService struct {
	reposPassengers *repository.Repository
}

func NewPassengersService(reposPassengers *repository.Repository) *PassengersService {
	return &PassengersService{reposPassengers}
}

func (p *PassengersService) CreatePassenger(routeId openapi_types.UUID, item models.PostPassenger) (models.Passenger, error) {

	return p.reposPassengers.CreatePassenger(routeId, item)
}
