package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"meteo-lightning/internal/domain/models"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type EnDB struct {
	log *slog.Logger
	db  *sql.DB
}

func NewEnDB(log *slog.Logger, dsn string) (*EnDB, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to open postgres db: %w", err)
	}

	return &EnDB{
		log: log,
		db:  db,
	}, nil

}

func (edb *EnDB) Close() {
	if edb.db != nil {
		edb.db.Close()
	}
}

func (edb *EnDB) SaveEnData(ctx context.Context, data []models.StrokeEN) error {

	fmt.Println("len data in storage:", len(data))
	tx, err := edb.db.Begin()
	if err != nil {
		return err
	}

	for _, el := range data {
		_, err := tx.ExecContext(ctx,
			`INSERT INTO enstrikes (cloud, time, latitude, longitude, signal, height, sensors)
				 VALUES($1,$2,$3,$4,$5,$6,$7)`,
			el.Cloud(), el.Time(), el.Lat(), el.Long(), el.Signal(), el.Height(), el.Sensors(),
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}
func (edb *EnDB) LightningData(ctx context.Context) ([]models.StrokeEN, error) {
	strokes := make([]models.StrokeEN, 0, 10000)

	return strokes, nil
}
