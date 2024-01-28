package science

import (
	"context"
	"fmt"
	"log/slog"
	"meteo-lightning/internal/config"
	"meteo-lightning/internal/domain/models"
	"meteo-lightning/internal/lib/logger/sl"
	"time"
)

type MeteoSource interface {
	// MeteoDataByTimeAndStation(ctx context.Context, t1, t2 time.Time, s models.Station) ([]models.MeteoData, error)
	StationMeteoParamsByTime(ctx context.Context, st models.Station, t time.Time, dur time.Duration) (models.MeteoParams, error)
	Close()
}
type StationsProvider interface {
	Stations(ctx context.Context) ([]models.Station, error)
	Close()
}

type LightningSource interface {
	StationLightningActivityByTime(ctx context.Context, st models.Station, t time.Time, dur time.Duration) ([]*models.StrokeEN, error)
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

func (s *ScienceService) MakeResearch(ctx context.Context) ([]*models.CorrPoint, error) {

	op := "science.MakeResearch"

	s.log.Info(op)

	cfg := config.MustLoadCfg()

	resCfg := cfg.Flags

	stations, err := s.stProv.Stations(ctx)
	if err != nil {
		return nil, err
	}

	// s.log.Info("st", slog.Any("stations", stations))

	points := make([]*models.CorrPoint, 0, 1000)

	for _, el := range stations {

		station := el

		begin := resCfg.Begin

		for begin.Before(resCfg.End) {
			point, err := models.NewCorrPoint(&station, resCfg.Dur)
			if err != nil {
				s.log.Error(op, sl.Err(err))
				continue
			}

			mParam, err := s.meteoProv.StationMeteoParamsByTime(ctx, station, begin, resCfg.Dur)
			if err != nil {
				s.log.Error(op, sl.Err(err))
				continue
			}

			strokes, err := s.strokeProv.StationLightningActivityByTime(ctx, station, begin, resCfg.Dur)
			if err != nil {
				s.log.Error(op, sl.Err(err))
				continue
			}

			point.SetMParams(&mParam)

			la := models.NewLActivity(strokes)

			point.SetlActivity(&la)

			points = append(points, point)

			begin = begin.Add(resCfg.Dur)

			break
		}
		break
	}

	fmt.Printf("%+v\n", points)

	return points, nil
}
