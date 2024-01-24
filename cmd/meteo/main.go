package main

import (
	"context"
	"fmt"
	"log/slog"
	"meteo-lightning/internal/config"
	"meteo-lightning/internal/filesource/meteofile"
	"meteo-lightning/internal/lib/logger/sl"
	"meteo-lightning/internal/services/meteoservice"
	"meteo-lightning/internal/storage/postgres"
	"time"
)

func main() {

	err := run()
	if err != nil {
		panic(err)
	}

}

func run() error {

	ctx := context.Background() //ctx

	cfg := config.MustLoadCfg() //load config

	log := sl.SetupLogger(cfg.Env) //log

	meteoDB, err := postgres.NewMeteoDB(log, cfg.DBconfig.ToString()) // DB
	if err != nil {
		return err
	}

	meteoSrv := meteoservice.NewService(meteoDB) // MeteoService

	files, err := meteofile.Files() // search files with meteo data
	if err != nil {
		return err
	}

	for _, el := range files {
		data, err := meteofile.Data(el)
		if err != nil {
			fmt.Printf("unable to read meteodata %v\n", err)
			continue
		}

		log.Info("readed data ", slog.String("from file", el))

		t := time.Now()
		err = meteoSrv.SaveMeteoData(ctx, data)
		if err != nil {
			log.Error("unable to save data", err)
			continue
		}
		log.Info("saved data ", slog.String("file", el), slog.Duration("time", time.Since(t)))
	}

	return nil
}
