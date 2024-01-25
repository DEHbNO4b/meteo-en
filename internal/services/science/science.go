package science

import (
	"context"
	"log/slog"
	"meteo-lightning/internal/domain/models"
	"time"
)

type MeteoSource interface {
	MeteoData(ctx context.Context) ([]models.MeteoData, error)
	// Stations(ctx context.Context) ([]models.Station, error)
	Close()
}

type LightningSource interface {
	LightningData(ctx context.Context) ([]models.StrokeEN, error)
	Close()
}

type ScienceConfiguration func(os *ScienceService) error

type ScienceService struct {
	log    *slog.Logger
	meteo  MeteoSource
	stroke LightningSource
}

func WithLogger(log *slog.Logger) ScienceConfiguration {
	return func(s *ScienceService) error {
		s.log = log
		return nil
	}
}

func New(ms MeteoSource, ls LightningSource, cfgs ...ScienceConfiguration) (*ScienceService, error) {

	s := &ScienceService{}
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

func (s *ScienceService) Close() {
	s.meteo.Close()
	s.stroke.Close()
}

func (s *ScienceService) MakeResearch(ctx context.Context) error {

	op := "science.MakeResearch"

	md, err := s.meteo.MeteoData(ctx)
	if err != nil {
		return err
	}
	// s.log.Info(op, slog.String("research", "success"))
	// l := len(md)
	// s.log.Info(op, slog.Int("data len", l))

	return nil
}

func (s *ScienceService) dataDurations(ctx context.Context, dur time.Duration, data []models.MeteoData, out chan<- models.MeteoData) error {

}
