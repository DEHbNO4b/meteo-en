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
	StationLightningActivityByTime(ctx context.Context, st models.Station, t time.Time, dur time.Duration) (models.LightningActivity, error)
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

	cfg := config.MustLoadCfg()
	resCfg := cfg.ResCfg

	stations, err := s.stProv.Stations(ctx)
	if err != nil {
		s.log.Error("err", sl.Err(err))
		return nil, fmt.Errorf("%s %w", op, err)
	}

	points := make([]*models.CorrPoint, 0, 1000)

	for _, el := range stations {

		begin := resCfg.Begin

		for begin.Before(resCfg.End) {
			point, err := models.NewCorrPoint(&el, resCfg.Dur)
			if err != nil {
				s.log.Error(op, sl.Err(err))
				continue
			}

			mParam, err := s.meteoProv.StationMeteoParamsByTime(ctx, el, begin, resCfg.Dur)
			if err != nil {
				s.log.Error(op, sl.Err(err))
				continue
			}

			lActivity, err := s.strokeProv.StationLightningActivityByTime(ctx, el, begin, resCfg.Dur)
			if err != nil {
				s.log.Error(op, sl.Err(err))
				continue
			}

			point.SetMParams(&mParam)
			point.SetlActivity(&lActivity)

			begin = begin.Add(resCfg.Dur)
			// begin = newTime
			points = append(points, point)
			break
		}
		break
	}

	return points, nil
}

// func (s *ScienceService) durationGen(ctx context.Context, dur time.Duration, data []models.MeteoData) <-chan models.MeteoData {

// 	out := make(chan models.MeteoData)
// 	// defer close(out)

// 	// mdata := models.MeteoData{}

// 	// var t time.Time

// 	// for _, el := range data {
// 	// 	if t.IsZero() {
// 	// 		t = el.Time.Add(dur)
// 	// 	}

// 	// }

// 	return out
// }
