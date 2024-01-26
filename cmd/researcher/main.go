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

	cfg := config.MustLoadCfg() //load config

	log := sl.SetupLogger(cfg.Env) //log

	mdb, err := postgres.NewMeteoDB(log, cfg.DBconfig.ToString()) // TODO: open meteo db
	if err != nil {
		return err
	}

	endb, err := postgres.NewEnDB(log, cfg.DBconfig.ToString()) // TODO: open en db
	if err != nil {
		return err
	}

	sdb, err := postgres.NewStationsDB(log, cfg.DBconfig.ToString()) // TODO: open stations db
	if err != nil {
		return err
	}

	srv, err := science.New(mdb, endb, sdb, science.WithLogger(log)) // TODO: create science service
	if err != nil {
		log.Error(op, err)
	}
	defer srv.Close()

	srv.MakeResearch(ctx) // TODO: make research

	// TODO: save results

	return nil
}
