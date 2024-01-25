package main

import (
	"context"
	"meteo-lightning/internal/config"
	"meteo-lightning/internal/lib/logger/sl"
	"meteo-lightning/internal/services/science"
	"meteo-lightning/internal/storage/postgres"
)

type MeteoProvider interface {
}

type LightningProvider interface {
}

func main() {

	err := run()
	if err != nil {
		panic(err)
	}

}

func run() error {
	op := "researcher.main.run"

	ctx := context.Background()

	//load config
	cfg := config.MustLoadCfg()

	log := sl.SetupLogger(cfg.Env) //log

	// TODO: open mete db
	mdb, err := postgres.NewMeteoDB(log, cfg.DBconfig.ToString())
	if err != nil {
		return err
	}

	endb, err := postgres.NewEnDB(log, cfg.DBconfig.ToString())
	if err != nil {
		return err
	}

	// TODO: create science service

	srv, err := science.New(mdb, endb, science.WithLogger(log))
	if err != nil {
		log.Error(op, err)
	}
	defer srv.Close()

	// TODO: make research
	srv.MakeResearch(ctx)

	// TODO: save results

	return nil
}
