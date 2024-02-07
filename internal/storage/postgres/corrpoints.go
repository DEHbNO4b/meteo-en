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

	lmp := domainMeteoParToLoc(cp.MeteoParams)

	_, err := cdb.db.ExecContext(ctx, `INSERT INTO corrpoints (station,wind_speed,maxwind_speed,rain,max_rain,rain_rate,maxrain_rate,
													count,maxpozitivesignal,maxnegativesignal,pozitivesignal,negativesignal,cloudtyperelation,groundtyperelation,abs_signal)
									VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)`,
		cp.Name(), lmp.WindSpeed, lmp.MaxWindSpeed, lmp.Rain, lmp.MaxRain, lmp.RainRate, lmp.MaxRainRate,
		cp.Count(), cp.MaxPozSig(), cp.MaxNegSig(), cp.PozSig(), cp.NegSig(), cp.CloudTypeRel(), cp.GroundTypeRel(), cp.AbsSig())
	if err != nil {
		return err
	}

	return nil
}

func (cdb *CorrpointDB) CorrParams(ctx context.Context) ([]models.CorrPoint, error) {

	op := "storage/postgres/corrpoint.CorrParams"
	ans := make([]models.CorrPoint, 0, 1000)

	rows, err := cdb.db.QueryContext(ctx, `SELECT wind_speed,maxwind_speed,rain,max_rain,rain_rate,maxrain_rate,
	count,maxpozitivesignal,maxnegativesignal,pozitivesignal,negativesignal,cloudtyperelation,groundtyperelation,abs_signal 
	FROM corrpoints  `)
	//where station like 'Chamlik' where station like 'Labinsk' or station like 'VGI'
	//where time between '
	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			cp              = models.CorrPoint{}
			mp              = models.MeteoParams{}
			la              = models.LightningActivity{}
			count           int
			maxPoz, maxNeg  int64
			pozSig, negSig  int64
			cloudT, groundT float64
			absSig          float64
		)
		cp.LightningActivity = &la
		cp.MeteoParams = &mp

		if err := rows.Scan(&cp.WindSpeed, &cp.HiSpeed, &cp.Rain, &cp.MaxRain, &cp.RainRate, &cp.MaxRainRate,
			&count, &maxPoz, &maxNeg, &pozSig, &negSig, &cloudT, &groundT, &absSig); err != nil {
			return nil, fmt.Errorf("%s %w", op, err)
		}

		cp.SetCount(count)
		cp.SetMaxPozitiveSignal(maxPoz)
		cp.SetMaxNegativeSignal(maxNeg)
		cp.SetPozSignal(pozSig)
		cp.SetNegSignal(negSig)
		cp.SetCloudTypeRelation(cloudT)
		cp.SetGroundTypeRelation(groundT)
		cp.SetAbsSig(absSig)

		ans = append(ans, cp)
	}

	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}

	return ans, nil

}
