package main

import (
	"context"
	"fmt"
	"meteo-lightning/internal/domain/models"
	"meteo-lightning/internal/filesource/meteofile"
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

	// //load config
	// cfg := config.MustLoadCfg()

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
		// fmt.Println(data)
		// break
	}

	// ms.Read("public/meteo/yspenskoe2023.txt")

	return nil
}
