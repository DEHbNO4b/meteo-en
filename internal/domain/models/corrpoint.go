package models

import (
	"errors"
	"time"
)

type CorrPoint struct {
	station   *Station
	dur       time.Duration
	lActivity *LightningActivity
	mParams   *MeteoParams
}

func NewCorrPoint(s *Station, dur time.Duration) (*CorrPoint, error) {
	if s == nil || dur.Seconds() < 60 {
		return nil, errors.New("wrong init value for corr point")
	}

	return &CorrPoint{
		station: s,
		dur:     dur,
	}, nil
}

func (c *CorrPoint) SetStation(s *Station) {
	c.station = s
}

func (c *CorrPoint) SetDur(d time.Duration) {
	c.dur = d
}

func (c *CorrPoint) SetlActivity(l *LightningActivity) {
	c.lActivity = l
}

func (c *CorrPoint) SetMParams(mp *MeteoParams) {
	c.mParams = mp
}

func (c *CorrPoint) Station() *Station {
	return c.station
}

func (c *CorrPoint) Dur() time.Duration {
	return c.dur
}

func (c *CorrPoint) LightningActivity() *LightningActivity {
	return c.lActivity
}

func (c *CorrPoint) MetParams() *MeteoParams {
	return c.mParams
}
