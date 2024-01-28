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

	ctx := context.Background()

	cfg := config.MustLoadCfg() //load config

	log := sl.SetupLogger(cfg.Env) //log

	mdb, err := postgres.NewMeteoDB(log, cfg.DBconfig.ToString()) // open meteo db
	if err != nil {
		return err
	}

	endb, err := postgres.NewEnDB(log, cfg.DBconfig.ToString()) // open en db
	if err != nil {
		return err
	}

	sdb, err := postgres.NewStationsDB(log, cfg.DBconfig.ToString()) // open stations db
	if err != nil {
		return err
	}

	srv, err := science.New(mdb, endb, sdb, science.WithLogger(log)) // create science service
	if err != nil {
		return err
	}
	defer srv.Close()

	_, err = srv.MakeResearch(ctx) // make research
	if err != nil {
		return err
	}

	// TODO: save results

	return nil
}
