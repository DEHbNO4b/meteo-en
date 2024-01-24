package main

import (
	"context"
	"fmt"
	"log/slog"
	"meteo-lightning/internal/config"
	"meteo-lightning/internal/filesource/enfile"
	"meteo-lightning/internal/lib/logger/sl"
	"meteo-lightning/internal/services/enservice"
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

	enDB, err := postgres.NewEnDB(log, cfg.DBconfig.ToString()) // DB
	if err != nil {
		return err
	}

	enSrv := enservice.NewService(enDB) // MeteoService

	files, err := enfile.Files() // search files with meteo data
	if err != nil {
		return err
	}

	for _, el := range files {

		data, err := enfile.Data(el)

		if err != nil {
			fmt.Printf("unable to read en data %v\n", err)
			continue
		}

		log.Info("readed data ", slog.String("from file", el))

		t := time.Now()
		err = enSrv.SaveEnData(ctx, data)
		if err != nil {
			log.Error("unable to save en data", err)
			continue
		}
		log.Info("saved en data ", slog.String("file", el), slog.Duration("time", time.Since(t)))
	}

	return nil
}
