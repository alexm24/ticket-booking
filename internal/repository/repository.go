package repository

import (
	"github.com/jmoiron/sqlx"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"

	"github.com/alexm24/ticket-booking/internal/models"
)

type IRoutes interface {
	CreateRoute(item models.PostRoute) (models.Route, error)
}

type IPassengers interface {
	CreatePassenger(routeId openapi_types.UUID, item models.PostPassenger) (models.Passenger, error)
}

type Repository struct {
	IRoutes
	IPassengers
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IRoutes:     NewRoutesPostgres(db),
		IPassengers: NewPassengersPostgres(db),
	}
}
