package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"meteo-lightning/internal/domain/models"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type RGMMeteoDB struct {
	log *slog.Logger
	db  *sql.DB
}

func NewRGMMeteoDB(log *slog.Logger, dsn string) (*RGMMeteoDB, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to open postgres db: %w", err)
	}

	return &RGMMeteoDB{
		log: log,
		db:  db,
	}, nil
}

func (mdb *RGMMeteoDB) Close() {
	if mdb.db != nil {
		mdb.db.Close()
	}
}

func (mdb *RGMMeteoDB) StationMeteoParamsByTime(ctx context.Context, st models.Station, t time.Time, dur time.Duration) (models.MeteoParams, error) {

	op := "storage/postgres/MeteoDataRGM.StationMeteoParamsByTime"

	row := mdb.db.QueryRowContext(ctx, `SELECT wind_speed,r,r24	FROM meteodata_rgm as m WHERE m.station_id =$1  AND time BETWEEN $2 AND $3;`,
		st.StationID(), t, t.Add(dur))

	var (
		windSpeed  string
		wavg, wmax int
		r, r24     float64
	)
	if err := row.Scan(&windSpeed, &r, &r24); err != nil {
		return models.MeteoParams{}, fmt.Errorf("%s %w", op, err)
	}

	if err := row.Err(); err != nil {
		return models.MeteoParams{}, fmt.Errorf("%s %w", op, err)
	}

	wavg, wmax, err := parseWind(windSpeed)
	if err != nil {
		return models.MeteoParams{}, err
	}

	mp := models.MeteoParams{}

	mp.WindSpeed = float64(wavg)
	mp.HiSpeed = float64(wmax)
	mp.MaxRain = r
	mp.MaxRainRate = r24

	return mp, nil
}

func (mdb *RGMMeteoDB) StationDataTimes(ctx context.Context, st models.Station) (time.Time, time.Time, error) {

	op := "storage/postgres/MeteoData.StationDataTimes"

	var t1, t2 time.Time

	row := mdb.db.QueryRowContext(ctx, `SELECT MAX(time),MIN(time)
	 								FROM meteodata_rgm as m WHERE m.station_id=$1  `,
		st.StationID())

	if err := row.Scan(&t1, &t2); err != nil {
		return t1, t2, fmt.Errorf("%s %w", op, err)
	}
	return t1, t2, nil
}
