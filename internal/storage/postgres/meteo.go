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
			`INSERT INTO meteodata (station, time, temp_out, wind_speed, wind_dir, wind_run, wind_chill, bar, rain,rain_rate)
				 VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`,
			el.StName, el.Time, el.TempOut, el.WindSpeed, el.WindDir, el.WindRun, el.WindChill, el.Bar, el.Rain, el.RainRate,
		)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("%s %w", op, err)
		}
	}

	tx.Commit()

	return nil
}

func (mdb *MeteoDB) StationMeteoParamsByTime(ctx context.Context, st models.Station, t time.Time, dur time.Duration) (*models.MeteoParams, error) {

	op := "storage/postgres/MeteoData.StationMeteoParamsByTime"

	mp := models.MeteoParams{}

	row := mdb.db.QueryRowContext(ctx, `SELECT count(*),AVG(wind_speed),AVG(rain),AVG(rain_rate),MAX(wind_speed),MAX(rain),MAX(rain_rate)
	 								FROM meteodata as m WHERE m.station ~ $1   AND time BETWEEN $2 AND $3`,
		st.Name(), t, t.Add(dur))

	var (
		count                              int
		windSpeed, rain, rainRate          sql.NullFloat64
		maxWindSpeed, maxRain, maxRainRate sql.NullFloat64
	)
	if err := row.Scan(&count, &windSpeed, &rain, &rainRate, &maxWindSpeed, &maxRain, &maxRainRate); err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
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
		mp.MaxWindSpeed = maxWindSpeed.Float64
	}
	if maxRain.Valid {
		mp.MaxRain = maxRain.Float64
	}
	if maxRainRate.Valid {
		mp.MaxRainRate = maxRainRate.Float64
	}

	if count == 0 {
		return nil, fmt.Errorf("%s %w", op, storage.ErrNoDataFound)
	}

	return &mp, nil
}

// func (mdb *MeteoDB) MeteoDataByTimeAndStation(ctx context.Context, t1, t2 time.Time, s models.Station) ([]models.MeteoData, error) {

// 	op := "storage/postgres/MeteoData"

// 	ans := make(map[string][]models.MeteoData)

// 	var data []models.MeteoData

// 	rows, err := mdb.db.Query(`SELECT m.station, m.time, m.temp_out, m.wind_speed, m.wind_dir,
// 								m.wind_run, m.wind_chill, m.bar, m.rain, m.rain_rate,
// 								s.name, s.lat, s.long
// 								FROM meteodata AS m
// 								JOIN stations AS s ON  m.station LIKE '%'||s.name||'%'
// 								ORDER BY s.name,m.time`)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, fmt.Errorf("%s %w", op, storage.ErrNoDataFound)
// 		}
// 		return nil, fmt.Errorf("%s %w", op, err)
// 	}

// 	defer rows.Close()

// 	for rows.Next() {

// 		var (
// 			name      string
// 			lat, long float64
// 		)
// 		md := models.MeteoData{}
// 		station := models.Station{}

// 		md.Station = station

// 		if err := rows.Scan(&md.StName, &md.Time, &md.TempOut, &md.WindSpeed, &md.WindDir,
// 			&md.WindRun, &md.WindChill, &md.Bar, &md.Rain, &md.RainRate,
// 			&name, &lat, &long,
// 		); err != nil {
// 			mdb.log.Error(op, sl.Err(err))
// 			continue
// 		}

// 		station.SetName(name)
// 		station.SetLat(lat)
// 		station.SetLong(long)

// 		md.Station = station

// 		d, ok := ans[md.Station.Name()]
// 		if !ok {
// 			data = make([]models.MeteoData, 0, 5000)
// 		} else {
// 			data = d
// 		}

// 		data = append(data, md)
// 		ans[md.Station.Name()] = data

// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return nil, nil
// }
