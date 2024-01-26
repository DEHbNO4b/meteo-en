package main

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

	// cfg := config.MustLoadCfg() // load config

	// TODO: open DB

	// TODO: create science service

	// TODO: make research

	// TODO: save results

	return nil
}
