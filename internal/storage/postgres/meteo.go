package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"meteo-lightning/internal/domain/models"
	"meteo-lightning/internal/storage"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
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

	op := "storage/postgres/MeteoData.SaveMeteoData"

	fmt.Println("len data in storage:", len(data))
	tx, err := mdb.db.Begin()
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	for _, el := range data {
		_, err := tx.ExecContext(ctx,
			`INSERT INTO meteodata (station, time, temp_out, wind_speed, wind_dir, wind_run,hi_speed, wind_chill, bar, rain,rain_rate)
				 VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
			el.StName, el.Time, el.TempOut, el.WindSpeed, el.WindDir, el.WindRun, el.HiSpeed, el.WindChill, el.Bar, el.Rain, el.RainRate,
		)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("%s %w", op, err)
		}
	}

	tx.Commit()

	return nil
}

func (mdb *MeteoDB) StationMeteoParamsByTime(ctx context.Context, st models.Station, t time.Time, dur time.Duration) (models.MeteoParams, error) {

	op := "storage/postgres/MeteoData.StationMeteoParamsByTime"

	mp := models.MeteoParams{}

	row := mdb.db.QueryRowContext(ctx, `SELECT count(*),AVG(wind_speed),AVG(rain),AVG(rain_rate),MAX(hi_speed),MAX(rain),MAX(rain_rate)
	 								FROM meteodata as m WHERE m.station ~ $1   AND time BETWEEN $2 AND $3`,
		st.Name(), t, t.Add(dur))

	var (
		count                              int
		windSpeed, rain, rainRate          sql.NullFloat64
		maxWindSpeed, maxRain, maxRainRate sql.NullFloat64
	)
	if err := row.Scan(&count, &windSpeed, &rain, &rainRate, &maxWindSpeed, &maxRain, &maxRainRate); err != nil {
		return mp, fmt.Errorf("%s %w", op, err)
	}

	if windSpeed.Valid {
		mp.WindSpeed = windSpeed.Float64
	}
	if rain.Valid {
		mp.Rain = rain.Float64
	}
	if rainRate.Valid {
		mp.RainRate = rainRate.Float64
	}
	if maxWindSpeed.Valid {
		mp.HiSpeed = maxWindSpeed.Float64
	}
	if maxRain.Valid {
		mp.MaxRain = maxRain.Float64
	}
	if maxRainRate.Valid {
		mp.MaxRainRate = maxRainRate.Float64
	}

	if count == 0 {
		return mp, fmt.Errorf("%s %w", op, storage.ErrNoDataFound)
	}

	return mp, nil
}

func (mdb *MeteoDB) StationDataTimes(ctx context.Context, st models.Station) (time.Time, time.Time, error) {

	op := "storage/postgres/MeteoData.StationDataTimes"

	var t1, t2 time.Time

	row := mdb.db.QueryRowContext(ctx, `SELECT MAX(time),MIN(time)
	 								FROM meteodata as m WHERE m.station ~ $1  `,
		st.Name())

	if err := row.Scan(&t1, &t2); err != nil {
		return t1, t2, fmt.Errorf("%s %w", op, err)
	}
	return t1, t2, nil
}
