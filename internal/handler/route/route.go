package route

import "github.com/alexm24/ticket-booking/internal/service"

type Route struct {
	service *service.Service
}

func NewRoute(service *service.Service) *Route {
	return &Route{service}
}
