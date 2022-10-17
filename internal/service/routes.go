package service

import (
	"github.com/alexm24/ticket-booking/internal/models"
	"github.com/alexm24/ticket-booking/internal/repository"
)

type RoutesService struct {
	reposRoutes *repository.Repository
}

func NewRoutesService(reposRoutes *repository.Repository) *RoutesService {
	return &RoutesService{reposRoutes}
}

func (r *RoutesService) CreateRoute(item models.PostRoute) (models.Route, error) {
	return r.reposRoutes.CreateRoute(item)
}
