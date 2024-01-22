package postgres

import (
	"database/sql"
	"errors"
	"log/slog"
)

var (
	ErrInvalidDB = errors.New("invalid db ")
)

type MeteoDB struct {
	log *slog.Logger
	db  *sql.DB
}

func NewMeteo(log *slog.Logger, db *sql.DB) (*MeteoDB, error) {

	if db == nil {
		return nil, ErrInvalidDB
	}

	return &MeteoDB{
		log: log,
		db:  db,
	}, nil

}
