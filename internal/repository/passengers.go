package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"

	"github.com/alexm24/ticket-booking/internal/models"
)

const (
	passengersTable = "passengers"
)

type PassengersPostgres struct {
	db *sqlx.DB
}

func NewPassengersPostgres(db *sqlx.DB) *PassengersPostgres {
	return &PassengersPostgres{db}
}

func (p *PassengersPostgres) CreatePassenger(routeId openapi_types.UUID, item models.PostPassenger) (models.Passenger, error) {
	var i models.Passenger
	q := fmt.Sprintf(
		`INSERT INTO %s
				(id, route_id, fullname, age, phone, email) values (uuid_generate_v4(), $1, $2, $3, $4, $5) RETURNING *;`,
		passengersTable)
	row := p.db.QueryRowx(q, routeId, *item.Fullname, *item.Age, *item.Phone, *item.Email)
	if err := row.StructScan(&i); err != nil {
		return models.Passenger{}, err
	}
	return i, nil
}
