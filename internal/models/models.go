package models

import "github.com/alexm24/ticket-booking/internal/handler/api"

type PostPassenger api.AddPassengerJSONBody

type PostRoute api.AddRouteJSONRequestBody

type Route struct {
	api.SIdentifier
	api.SRoute
}

type Passenger struct {
	api.SIdentifier
	api.SPassenger
}
