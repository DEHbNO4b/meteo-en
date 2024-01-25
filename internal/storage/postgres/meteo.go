package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"meteo-lightning/internal/domain/models"
	"meteo-lightning/internal/lib/logger/sl"
	"meteo-lightning/internal/storage"

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
func (mdb *MeteoDB) Close() {
	if mdb.db != nil {
		mdb.db.Close()
	}
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
			el.StName, el.Time, el.TempOut, el.WindSpeed, el.WindDir, el.WindRun, el.WindChill, el.Bar, el.Rain, el.RainRate,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

func (mdb *MeteoDB) MeteoData(ctx context.Context) ([]models.MeteoData, error) {

	op := "storage/postgres/MeteoData"

	data := make([]models.MeteoData, 0, 1000)

	rows, err := mdb.db.Query(`SELECT m.station, m.time, m.temp_out, m.wind_speed, m.wind_dir, 
								m.wind_run, m.wind_chill, m.bar, m.rain, m.rain_rate,
								s.name, s.lat, s.long 
								FROM meteodata AS m
								JOIN stations AS s ON  m.station LIKE '%'||s.name||'%'
								ORDER BY s.name,m.time`)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return data, fmt.Errorf("%s %w", op, storage.ErrNoDataFound)
		}
		return data, fmt.Errorf("%s %w", op, err)
	}

	defer rows.Close()

	for rows.Next() {

		var (
			name      string
			lat, long float64
		)
		md := models.MeteoData{}
		station := models.Station{}

		md.Station = station

		if err := rows.Scan(&md.StName, &md.Time, &md.TempOut, &md.WindSpeed, &md.WindDir,
			&md.WindRun, &md.WindChill, &md.Bar, &md.Rain, &md.RainRate,
			&name, &lat, &long,
		); err != nil {
			mdb.log.Error(op, sl.Err(err))
			continue
		}

		station.SetName(name)
		station.SetLat(lat)
		station.SetLong(long)

		md.Station = station

		data = append(data, md)
	}

	err = rows.Err()
	if err != nil {
		return data, err
	}

	return data, nil
}
