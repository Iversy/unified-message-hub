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
	err = storage.initTables()
	if err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *PGstorage) initTables() error {
	sql := fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %v (
        %v SERIAL PRIMARY KEY,
        %v VARCHAR(100) NOT NULL,
        %v VARCHAR(255) UNIQUE NOT NULL,
        %v INT
    )`, tableName, IDСolumnName, NameСolumnName, EmailСolumnName, AgeСolumnName)
	_, err := s.db.Exec(context.Background(), sql)
	if err != nil {
		return err
	}
	return nil
}
