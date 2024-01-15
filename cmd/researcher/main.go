package main

import (
	"fmt"
	"meteo-lightning/internal/source/meteo"
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

	// Create meteo file sorce struct
	s, err := meteo.New("./public/meteo", "public/meteo/*.txt")
	if err != nil {
		return err
	}

	// search files with meteo data
	files, err := s.Search()
	if err != nil {
		return err
	}

	fmt.Println(files)

	return nil
}
