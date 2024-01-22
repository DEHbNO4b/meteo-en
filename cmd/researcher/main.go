package main

import (
	"meteo-lightning/internal/config"
	"meteo-lightning/internal/filesource/meteofile"
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

	//load config
	cfg := config.MustLoadCfg()

	// Create meteo file sorce struct
	ms, err := meteofile.New(cfg.Fcfg.MeteoPath, cfg.Fcfg.MeteoTemplate)
	if err != nil {
		return err
	}

	// search files with meteo data
	files, err := ms.Search()
	if err != nil {
		return err
	}

	// fmt.Printf("%+v\n", files)

	// s.Read("public/meteo/yspenskoe2023.txt")

	return nil
}
