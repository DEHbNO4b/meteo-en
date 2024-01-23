package main

import (
	"context"
	"fmt"
	"meteo-lightning/internal/config"
	"meteo-lightning/internal/domain/models"
	"meteo-lightning/internal/filesource/meteofile"
	"meteo-lightning/internal/lib/logger/sl"
	"meteo-lightning/internal/services/meteoservice"
	"meteo-lightning/internal/storage/postgres"
)

type MeteoStore interface {
	SaveMeteoData(ctx context.Context, data []models.MeteoData) error
}

// type LightningProvider interface {
// }

func main() {

	err := run()
	if err != nil {
		panic(err)
	}

}

func run() error {

	//load config
	cfg := config.MustLoadCfg()

	//log
	log := sl.SetupLogger(cfg.Env)

	// DB
	meteoDB, err := postgres.NewMeteoDB(log, cfg.DBconfig.ToString())
	if err != nil {
		return err
	}

	// MeteoService
	meteoservice.NewService(meteoDB)

	// search files with meteo data
	files, err := meteofile.Files()
	if err != nil {
		return err
	}
	for _, el := range files {
		data, err := meteofile.Data(el)
		if err != nil {
			fmt.Printf("unable to read meteodata %v\n", err)
		}

	}

	return nil
}
