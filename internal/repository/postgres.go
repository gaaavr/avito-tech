package repository

import (
	"avito/configs"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

// NewPostgresDB returns the pool for connecting to the database
func NewPostgresDB(cfg configs.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cfg.DbHost, cfg.DbPort, cfg.DbUsername, cfg.DbName, cfg.DbPassword, cfg.DbSslmode)
	dbpool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	if err = dbpool.Ping(context.Background()); err != nil {
		return nil, err
	}
	return dbpool, nil
}
