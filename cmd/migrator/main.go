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
		"postgres://"+cfg.DBconfig.User+":"+cfg.DBconfig.Password+"@"+cfg.DBconfig.Host+":"+cfg.DBconfig.Port+"/"+cfg.DBconfig.Database+"?sslmode=disable",
	)
	if err != nil {
		panic(err)
	}
	t, err := migrate.New(
		"file://"+migrationsPath,
		"postgres://test:test@localhost:5432/test?sslmode=disable",
	)
	if err != nil {
		panic(err)
	}

	// if err := m.Down(); err != nil {
	// 	panic(err)
	// }

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("osya: no migration change")

		}
		// panic(err)
	}

	if err := t.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("test: no migration change")
			return
		}
		panic(err)
	}

	fmt.Println("migrations applyed")

}
