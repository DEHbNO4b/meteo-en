package meteo

import "meteo-lightning/internal/source"

type MeteoSource struct {
	path string
}

func New(p string) (MeteoSource, error) {

	if p == "" {
		return MeteoSource{}, source.EmptyDataSource
	}

	return MeteoSource{path: p}, nil

}
