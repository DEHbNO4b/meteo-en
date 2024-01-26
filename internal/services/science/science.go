package science

import (
	"context"
	"fmt"
	"log/slog"
	"meteo-lightning/internal/domain/models"
	"meteo-lightning/internal/lib/logger/sl"
	"time"
)

type MeteoSource interface {
	MeteoDataByTimeAndStation(ctx context.Context, t1, t2 time.Time, s models.Station) ([]models.MeteoData, error)
	Close()
}
type StationsProvider interface {
	Stations(ctx context.Context) ([]models.Station, error)
	Close()
}

type LightningSource interface {
	LightningDataByTimeAndPos(ctx context.Context) ([]models.StrokeEN, error)
	Close()
}

type ScienceConfiguration func(os *ScienceService) error

type ScienceService struct {
	log        *slog.Logger
	meteoProv  MeteoSource
	strokeProv LightningSource
	stProv     StationsProvider
}

func WithLogger(log *slog.Logger) ScienceConfiguration {
	return func(s *ScienceService) error {
		s.log = log
		return nil
	}
}

func New(ms MeteoSource, ls LightningSource, sp StationsProvider, cfgs ...ScienceConfiguration) (*ScienceService, error) {

	s := &ScienceService{}

	s.meteoProv = ms
	s.strokeProv = ls
	s.stProv = sp

	for _, cfg := range cfgs {
		err := cfg(s)
		if err != nil {
			return s, err
		}
	}

	return s, nil
}

func (s *ScienceService) Close() {
	s.meteoProv.Close()
	s.strokeProv.Close()
	s.stProv.Close()
}

func (s *ScienceService) MakeResearch(ctx context.Context) error {

	op := "science.MakeResearch"

	// md, err := s.meteoProv.MeteoData(ctx)
	// if err != nil {
	// 	return err
	// }

	stations, err := s.stProv.Stations(ctx)
	if err != nil {
		s.log.Error("err", sl.Err(err))
		return fmt.Errorf("%s %w", op, err)
	}
	s.log.Info(op, slog.Any("stations", stations))

	// s.log.Info(op, slog.String("research", "success"))
	// l := len(md)
	// s.log.Info(op, slog.Int("data len", l))

	return nil
}

func (s *ScienceService) durationGen(ctx context.Context, dur time.Duration, data []models.MeteoData) <-chan models.MeteoData {

	out := make(chan models.MeteoData)
	// defer close(out)

	// mdata := models.MeteoData{}

	// var t time.Time

	// for _, el := range data {
	// 	if t.IsZero() {
	// 		t = el.Time.Add(dur)
	// 	}

	// }

	return out
}
