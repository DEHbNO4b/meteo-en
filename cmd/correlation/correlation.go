package correlation

import (
	"context"
	"meteo-lightning/internal/config"
	"meteo-lightning/internal/lib/logger/sl"
	"meteo-lightning/internal/services/science"
	"meteo-lightning/internal/storage/postgres"
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

	err = srv.MakeResearch(ctx) // make research
	if err != nil {
		return err
	}

	// TODO: save results

	return nil
}
