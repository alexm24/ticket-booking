package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/alexm24/ticket-booking/internal/models"
)

type RoutesPostgres struct {
	db *sqlx.DB
}

func NewRoutesPostgres(db *sqlx.DB) *RoutesPostgres {
	return &RoutesPostgres{db}
}

func (r *RoutesPostgres) CreateRoute(item models.PostRoute) (models.Route, error) {
	return models.Route{}, nil
}
