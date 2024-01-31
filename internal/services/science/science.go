package science

import (
	"context"
	"fmt"
	"log/slog"
	"meteo-lightning/internal/config"
	"meteo-lightning/internal/domain/models"
	"meteo-lightning/internal/lib/logger/sl"
	"meteo-lightning/internal/lib/semaphore"
	"sync"
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
type CorrPointProvider interface {
	SaveCorrpoint(ctx context.Context, cp models.CorrPoint) error
}

type ScienceConfiguration func(os *ScienceService) error

type ScienceService struct {
	log           *slog.Logger
	meteoProv     MeteoSource
	strokeProv    LightningSource
	stProv        StationsProvider
	corrPointProv CorrPointProvider
}

func WithLogger(log *slog.Logger) ScienceConfiguration {
	return func(s *ScienceService) error {
		s.log = log
		return nil
	}
}

func New(ms MeteoSource,
	ls LightningSource,
	sp StationsProvider,
	cpp CorrPointProvider,
	cfgs ...ScienceConfiguration) (*ScienceService, error) {

	s := &ScienceService{}

	s.meteoProv = ms
	s.strokeProv = ls
	s.stProv = sp
	s.corrPointProv = cpp

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

	cfg := config.MustLoadCfg()

	resCfg := cfg.Flags

	semaphore := semaphore.NewSemaphore(10)

	stations, err := s.stProv.Stations(ctx)
	if err != nil {
		return err
	}

	// points := make([]*models.CorrPoint, 0, 1000)
	t := time.Now()

	for _, el := range stations {

		station := el
		// _ = el
		s.log.Info("", slog.Any("station", el))
		begin := resCfg.Begin

		wg := sync.WaitGroup{}

		for i := 0; begin.Add(time.Duration(resCfg.Dur.Nanoseconds() * int64(i))).Before(resCfg.End); i++ {

			locI := i

			wg.Add(1)
			semaphore.Acquire()

			// fmt.Println("iteration:", locI)
			// fmt.Println(begin.Add(time.Duration(resCfg.Dur.Nanoseconds() * int64(locI))))

			go func() {

				defer wg.Done()
				defer semaphore.Release()

				point, err := models.NewCorrPoint(&station, resCfg.Dur)
				if err != nil {
					s.log.Error(op, sl.Err(err))
					// continue
					return
				}

				ldur := resCfg.Dur.Nanoseconds() * int64(locI)

				locT := begin
				locT.Add(time.Duration(ldur))
				// fmt.Println(locT)

				mParam, err := s.meteoProv.StationMeteoParamsByTime(ctx, station, locT, resCfg.Dur)
				if err != nil {
					s.log.Error(op, sl.Err(err))
					// continue
					return
				}
				if mParam.MaxWindSpeed != 0 {
					fmt.Printf("%+v\n", mParam.MaxWindSpeed)
				}

				strokes, err := s.strokeProv.StationLightningActivityByTime(ctx, station, begin, resCfg.Dur)
				if err != nil {
					s.log.Error(op, sl.Err(err))
					// continue
					return
				}

				// point.SetMParams(&mParam)
				point.MeteoParams = &mParam

				la := models.NewLActivity(strokes)

				// point.SetlActivity(&la)
				point.LightningActivity = &la

				// points = append(points, point)

				err = s.corrPointProv.SaveCorrpoint(ctx, *point)
				if err != nil {
					s.log.Error("unable to save corrpoint", sl.Err(err))
				}

				begin = begin.Add(resCfg.Dur)

			}()

		}

		wg.Wait()

	}

	fmt.Printf("research took %v", time.Since(t))

	return nil
}

func (s *ScienceService) CalculateCorr() ([]string, error) {
	ans := make([]string, 0, 16)

	return ans, nil

}
