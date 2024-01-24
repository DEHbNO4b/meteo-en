package main

import (
	"errors"
	"flag"
	"fmt"
	"meteo-lightning/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var (
		migrationsPath string
	)

	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.Parse()

	if migrationsPath == "" {
		panic("migrations path is required")
	}

	cfg := config.MustLoadCfg()

	m, err := migrate.New(
		"file://"+migrationsPath,
		// "postgres://practicum:practicum@localhost:5432/practicum?sslmode=disable",
		cfg.DBconfig.ToString(),
	)
	if err != nil {
		panic(err)
		// fmt.Println("err", err)
	}

	if err := m.Down(); err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migration change")
			return
		}
		panic(err)
	}
	fmt.Println("migrations applyed")

}
