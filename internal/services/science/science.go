package science

import (
	"log/slog"
	"meteo-lightning/internal/domain/models"
)

type MeteoSource interface {
	MeteoData() ([]models.MeteoData, error)
}

type LightningSource interface {
	LightningData() ([]models.StrokeEN, error)
}

type ScienceConfiguration func(os *Science) error

type Science struct {
	log    *slog.Logger
	meteo  MeteoSource
	stroke LightningSource
}

func WithLogger(log *slog.Logger) ScienceConfiguration {
	return func(s *Science) error {
		s.log = log
		return nil
	}
}

func New(ms MeteoSource, ls LightningSource, cfgs ...ScienceConfiguration) (*Science, error) {

	s := &Science{}
	s.meteo = ms
	s.stroke = ls

	for _, cfg := range cfgs {
		err := cfg(s)
		if err != nil {
			return s, err
		}
	}

	return s, nil
}
