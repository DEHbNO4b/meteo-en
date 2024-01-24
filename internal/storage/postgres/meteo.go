package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"meteo-lightning/internal/domain/models"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	ErrInvalidDB = errors.New("invalid db ")
)

type MeteoDB struct {
	log *slog.Logger
	db  *sql.DB
}

func NewMeteoDB(log *slog.Logger, dsn string) (*MeteoDB, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to open postgres db: %w", err)
	}

	return &MeteoDB{
		log: log,
		db:  db,
	}, nil

}

func (mdb *MeteoDB) SaveMeteoData(ctx context.Context, data []models.MeteoData) error {
	tx, err := mdb.db.BeginTx(ctx, nil)
	for _, el := range data {
		_, err := tx.ExecContext(ctx,
			"INSERT INTO ")
	}
}
