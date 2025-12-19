package pgstorage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGstorage struct {
	db *pgxpool.Pool
}

func NewPGStorge(connString string) (*PGstorage, error) {

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}
	storage := &PGstorage{
		db: db,
	}
	return storage, nil
}
