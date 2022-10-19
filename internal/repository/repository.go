package repository

import (
	"avito/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

// User - Interface describing the user entity
type User interface {
	AccrualFunds(ac models.AccrualCash) error
}

// Repository - object responsible for the work of logic with the database
type Repository struct {
	User
}

// NewRepository - constructor function for Repository
func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		User: NewUserRepo(db),
	}
}
