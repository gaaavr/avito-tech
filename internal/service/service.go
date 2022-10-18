package service

import (
	"avito/internal/models"
	"avito/internal/repository"
)

// User - Interface describing the user entity
type User interface {
	AccrualFunds(ac models.AccrualCash) (code int, err error)
}

// Service - object responsible for the operation of the internal logic
type Service struct {
	User
}

// NewService - constructor function for Service
func NewService(repository *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repository.User),
	}
}
