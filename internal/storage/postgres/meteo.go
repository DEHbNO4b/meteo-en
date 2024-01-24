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

	fmt.Println("len data in storage:", len(data))
	tx, err := mdb.db.Begin()
	if err != nil {
		return err
	}

	for _, el := range data {
		_, err := tx.ExecContext(ctx,
			`INSERT INTO meteodata (station, time, temp_out, wind_speed, wind_dir, wind_run, wind_chill, bar, rain,rain_rate)
				 VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`,
			el.Station, el.Time, el.TempOut, el.WindSpeed, el.WindDir, el.WindRun, el.WindChill, el.Bar, el.Rain, el.RainRate,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}
