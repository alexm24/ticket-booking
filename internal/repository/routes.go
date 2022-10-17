package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	"github.com/alexm24/ticket-booking/internal/models"
)

const (
	routesTable = "routes"
)

type RoutesPostgres struct {
	db *sqlx.DB
}

func NewRoutesPostgres(db *sqlx.DB) *RoutesPostgres {
	return &RoutesPostgres{db}
}

func (r *RoutesPostgres) CreateRoute(item models.PostRoute) (models.Route, error) {
	var i models.Route
	query := fmt.Sprintf(`INSERT INTO %s (id, route) values (uuid_generate_v4(), $1) RETURNING *;`, routesTable)
	row := r.db.QueryRowx(query, *item.Route)

	if err := row.StructScan(&i); err != nil {
		return models.Route{}, err
	}
	return i, nil
}
