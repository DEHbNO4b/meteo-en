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

	//load config
	// cfg := config.MustLoadCfg()

	return nil
}
