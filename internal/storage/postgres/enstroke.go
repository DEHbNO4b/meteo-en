package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"meteo-lightning/internal/config"
	"meteo-lightning/internal/domain/models"
	"meteo-lightning/internal/storage"
	"time"

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
func (edb *EnDB) StationLightningActivityByTime(ctx context.Context, st models.Station, t time.Time, dur time.Duration) ([]*models.StrokeEN, error) {

	op := "storage/postgres/enstroke.StationLightningActivityByTime"

	cfg := config.MustLoadCfg()

	// la := models.LightningActivity{}
	strokes := make([]*models.StrokeEN, 0, 100)

	sq := st.CalculateSquareCorners(cfg.Flags.Radius)

	// fmt.Println("in LA query ", sq.LowerLeft.Lat(), sq.UpperLeft.Lat(), sq.UpperLeft.Long(), sq.UpperRight.Long(), t, t.Add(dur))

	rows, err := edb.db.QueryContext(ctx, `SELECT cloud,latitude,longitude,signal,height from enstrikes 
									WHERE latitude BETWEEN $1 AND $2  AND longitude BETWEEN $3 AND $4
									AND time BETWEEN $5 AND $6`,
		sq.LowerLeft.Lat(), sq.UpperLeft.Lat(), sq.UpperLeft.Long(), sq.UpperRight.Long(), t, t.Add(dur))
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, storage.ErrNoDataFound
		}

		return nil, fmt.Errorf("%s %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			lat, long      float64
			cloud          bool
			signal, height int
			stroke         models.StrokeEN
		)
		if err := rows.Scan(&cloud, &lat, &long, &signal, &height); err != nil {
			return nil, fmt.Errorf("%s %w", op, err)
		}

		stroke.SetCloud(cloud)
		stroke.SetLat(lat)
		stroke.SetLong(long)
		stroke.SetSignal(signal)
		stroke.SetHeight(height)

		strokes = append(strokes, &stroke)
	}

	// la.Strokes = strokes

	return strokes, nil

}
