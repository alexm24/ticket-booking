package repository

import (
	"github.com/jmoiron/sqlx"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"

	"github.com/alexm24/ticket-booking/internal/models"
)

type PassengersPostgres struct {
	db *sqlx.DB
}

func NewPassengersPostgres(db *sqlx.DB) *PassengersPostgres {
	return &PassengersPostgres{db}
}

func (p *PassengersPostgres) CreatePassenger(routeId openapi_types.UUID, item models.PostPassenger) (models.Passenger, error) {
	return models.Passenger{}, nil
}
