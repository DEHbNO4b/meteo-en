package science

import (
	"context"
	"fmt"
	"log/slog"
	"meteo-lightning/internal/config"
	"meteo-lightning/internal/domain/models"
	"meteo-lightning/internal/lib/logger/sl"
	"meteo-lightning/internal/lib/semaphore"
	"sort"
	"sync"
	"time"
)

type MeteoSource interface {
	// MeteoDataByTimeAndStation(ctx context.Context, t1, t2 time.Time, s models.Station) ([]models.MeteoData, error)
	StationMeteoParamsByTime(ctx context.Context, st models.Station, t time.Time, dur time.Duration) (*models.MeteoParams, error)
	StationDataTimes(ctx context.Context, st models.Station) (time.Time, time.Time, error)
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
	CorrParams(ctx context.Context) ([]models.CorrPoint, error)
	Close()
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
	s.corrPointProv.Close()
}

func (s *ScienceService) MakeResearch(ctx context.Context) error {

	op := "science.MakeResearch"

	cfg := config.MustLoadCfg()

	resCfg := cfg.Flags

	semaphore := semaphore.NewSemaphore(5)

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

		max, min, err := s.meteoProv.StationDataTimes(ctx, station)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if min.After(resCfg.End) || max.Before(begin) {
			continue
		}

		wg := sync.WaitGroup{}

		for i := 0; begin.Add(time.Duration(resCfg.Dur.Nanoseconds() * int64(i))).Before(resCfg.End); i++ {

			locI := i

			wg.Add(1)
			semaphore.Acquire()

			go func() {

				defer wg.Done()
				defer semaphore.Release()

				point, err := models.NewCorrPoint(&station, resCfg.Dur)
				if err != nil {
					s.log.Error(op, sl.Err(err))
					return
				}

				ldur := resCfg.Dur.Nanoseconds() * int64(locI)

				locT := begin
				locT = locT.Add(time.Duration(ldur))

				mParam, err := s.meteoProv.StationMeteoParamsByTime(ctx, station, locT, resCfg.Dur)
				if err != nil {
					// s.log.Error(op, sl.Err(err))s
					return
				}

				strokes, err := s.strokeProv.StationLightningActivityByTime(ctx, station, locT, resCfg.Dur)
				if err != nil {
					// s.log.Error(op, sl.Err(err))
					// return
				}

				if len(strokes) == 0 && mParam == nil {
					return
				}
				// s.log.Info("succes calc corrpoint")
				point.MeteoParams = mParam

				la := models.NewLActivity(strokes)

				point.LightningActivity = &la

				err = s.corrPointProv.SaveCorrpoint(ctx, *point)
				if err != nil {
					s.log.Error("unable to save corrpoint", sl.Err(err))
				}

				// begin = begin.Add(resCfg.Dur)

			}()

		}

		wg.Wait()

	}

	fmt.Printf("research took %v", time.Since(t))

	return nil
}

func (s *ScienceService) CalculateCorr(ctx context.Context) ([]string, error) {

	ans := make([]string, 0, 16)
	pairs := make([]*models.Pair, 0, 10)

	cors, err := s.corrPointProv.CorrParams(ctx)
	if err != nil {
		return nil, err
	}

	ws_count := models.NewPair("wind speed - lightning count")
	rr_count := models.NewPair("rain rait- lightning count")
	r_count := models.NewPair("rain - lightning count")

	ws_absSig := models.NewPair("wind speed - absolut signal")
	rr_absSig := models.NewPair("rain rait- absolut signal")
	r_absSig := models.NewPair("rain - absolut signal")

	ws_maxNeg := models.NewPair("wind speed - max negaqtive signal")
	rr_maxNeg := models.NewPair("rain rait- max negaqtive signal")
	r_maxNeg := models.NewPair("rain - max negaqtive signal")

	pairs = append(pairs, ws_count)
	pairs = append(pairs, rr_count)
	pairs = append(pairs, r_count)
	pairs = append(pairs, ws_absSig)
	pairs = append(pairs, rr_absSig)
	pairs = append(pairs, r_absSig)
	pairs = append(pairs, ws_maxNeg)
	pairs = append(pairs, rr_maxNeg)
	pairs = append(pairs, r_maxNeg)

	for _, cp := range cors {

		// if cp.MaxRain > 0.1 {
		r_count.AddPair(cp.MaxRain, float64(cp.Count()))
		r_absSig.AddPair(cp.MaxRain, cp.AbsSig())
		r_maxNeg.AddPair(cp.MaxRain, float64(cp.MaxNegSig()))
		// }

		if cp.HiSpeed > 2.0 {
			ws_count.AddPair(cp.HiSpeed, float64(cp.Count()))
			ws_absSig.AddPair(cp.HiSpeed, cp.AbsSig())
			ws_maxNeg.AddPair(cp.HiSpeed, float64(cp.MaxNegSig()))
		}
		// if cp.MaxRainRate > 3 {
		rr_count.AddPair(cp.MaxRainRate, float64(cp.Count()))
		rr_absSig.AddPair(cp.MaxRainRate, cp.AbsSig())
		rr_maxNeg.AddPair(cp.MaxRainRate, float64(cp.MaxNegSig()))
		// }

	}

	for _, el := range pairs {
		el.Calculate()
	}

	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].CorrCoef() < pairs[j].CorrCoef()
	})

	for _, el := range pairs {
		ans = append(ans, el.String())
	}

	return ans, nil
}
