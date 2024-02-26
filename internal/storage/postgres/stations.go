package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"meteo-lightning/internal/domain/models"
	"meteo-lightning/internal/lib/logger/sl"
	"meteo-lightning/internal/storage"
)

type StationsDB struct {
	log *slog.Logger
	db  *sql.DB
}

func NewStationsDB(log *slog.Logger, dsn string) (*StationsDB, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to open postgres db: %w", err)
	}

	return &StationsDB{
		log: log,
		db:  db,
	}, nil

}
func (sdb *StationsDB) Close() {
	if sdb.db != nil {
		sdb.db.Close()
	}
}
func (sdb *StationsDB) Stations(ctx context.Context) ([]models.Station, error) {

	op := "stations.LightningDataByTimeAndPos"

	stations := make([]models.Station, 0, 7)

	rows, err := sdb.db.QueryContext(ctx, `select name,lat,long,station_id from stations where  type ='rosgidromet' and lat > 42 and lat < 47 and long > 37 and  long<46;`)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.ErrNoDataFound
		}
		return nil, fmt.Errorf("%s %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {

		s := models.Station{}

		var (
			name      string
			lat, long float64
			id        int64
		)

		if err := rows.Scan(&name, &lat, &long, &id); err != nil {
			sdb.log.Error(op, sl.Err(err))
		}

		s.SetName(name)
		s.SetLat(lat)
		s.SetLong(long)
		s.SetStationID(id)

		stations = append(stations, s)
	}

	return stations, nil
}
