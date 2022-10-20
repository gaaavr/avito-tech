package repository

import (
	"avito/configs"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

// NewPostgresDB returns the pool for connecting to the database
func NewPostgresDB(cfg configs.ConfigDB) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cfg.DbHost, cfg.DbPort, cfg.DbUsername, cfg.DbName, cfg.DbPassword, cfg.DbSslmode)
	dbPool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	if err = dbPool.Ping(context.Background()); err != nil {
		return nil, err
	}
	return dbPool, nil
}
