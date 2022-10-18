package repository

import (
	"avito/internal/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

// UserRepo - user object in the repository layer
type UserRepo struct {
	db *pgxpool.Pool
}

// NewUserRepo - constructor function for UserRepo
func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{db: db}
}

// AccrualCash - method of accruing cash to the balance
func (u *UserRepo) AccrualFunds(ac models.AccrualCash) (code int, err error) {
	rows, err := u.db.Query(context.Background(), "select * from users ")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var data models.AccrualCash
		err = rows.Scan(&data.ID, &data.Amount)
		if err != nil {
			return 500, err
		}
	}
	return 200, nil
}
