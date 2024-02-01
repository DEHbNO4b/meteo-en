package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"meteo-lightning/internal/domain/models"
)

type CorrpointDB struct {
	log *slog.Logger
	db  *sql.DB
}

func NewCorrpointDB(log *slog.Logger, dsn string) (*CorrpointDB, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to open corrpoint db: %w", err)
	}

	return &CorrpointDB{
		log: log,
		db:  db,
	}, nil

}
func (cdb *CorrpointDB) Close() {
	if cdb.db != nil {
		cdb.db.Close()
	}
}

func (cdb *CorrpointDB) SaveCorrpoint(ctx context.Context, cp models.CorrPoint) error {

	_, err := cdb.db.ExecContext(ctx, `INSERT INTO corrpoints (station,wind_speed,maxwind_speed,rain,max_rain,rain_rate,maxrain_rate,
													count,maxpozitivesignal,maxnegativesignal,pozitivesignal,negativesignal,cloudtyperelation,groundtyperelation,abs_signal)
									VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)`,
		cp.Name(), cp.WindSpeed, cp.MaxWindSpeed, cp.Rain, cp.MaxRain, cp.RainRate, cp.MaxRainRate,
		cp.Count(), cp.MaxPozSig(), cp.MaxNegSig(), cp.PozSig(), cp.NegSig(), cp.CloudTypeRel(), cp.GroundTypeRel(), cp.AbsSig())
	if err != nil {
		return err
	}

	return nil
}

// func (cdb *CorrpointDB) WindSpeedLaCount(ctx context.Context) ([]float64, []float64, error) {

// 	rows, err := cdb.db.QueryContext(ctx, `SELECT wind_speed,maxwind_speed,count,abs_signal FROM corrpoints WHERE count<>0 OR wind_speed >1 `)
// }