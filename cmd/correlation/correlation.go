package correlation

import (
	"context"
	"log/slog"
	"meteo-lightning/internal/config"
	"meteo-lightning/internal/lib/logger/sl"
	"meteo-lightning/internal/services/science"
	"meteo-lightning/internal/storage/postgres"
	"strconv"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {

	ctx := context.Background()

	cfg := config.MustLoadCfg() //load config

	log := sl.SetupLogger(cfg.Env) //log

	cpdb, err := postgres.NewCorrpointDB(log, cfg.DBconfig.ToString()) // open stations db
	if err != nil {
		return err
	}

	srv, err := science.New(nil, nil, nil, cpdb, science.WithLogger(log)) // create science service
	if err != nil {
		return err
	}
	defer srv.Close()

	corr, err := srv.CalculateCorr(ctx) // make research
	if err != nil {
		return err
	}
	for i, el := range corr {
		log.Info("correlation", slog.String(strconv.Itoa(i), el))
	}

	// TODO: save results

	return nil
}
