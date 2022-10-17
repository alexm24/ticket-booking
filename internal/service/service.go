package service

import "github.com/alexm24/ticket-booking/internal/repository"

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
