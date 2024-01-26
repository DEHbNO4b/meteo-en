package models

import (
	"errors"
	"time"
)

type MeteoData struct {

	//
	Dur     time.Duration
	Station Station

	StName       string
	Time         time.Time
	TempOut      float64
	HiTemp       float64
	LowTemp      float64
	OutHum       int
	DewPt        float64
	WindSpeed    float64
	WindDir      string
	WindRun      float64
	HiSpeed      float64
	HiDir        string
	WindChill    float64
	HeatIndex    float64
	THWIndex     float64
	Bar          float64
	Rain         float64
	RainRate     float64
	HeatD_D      float64
	CoolD_D      float64
	InTemp       float64
	InHum        float64
	InDew        float64
	InHeat       float64
	InEMC        float64
	InAirDensity float64
	WindSamp     float64
	WindTx       float64
	ISSRecept    float64
	ArcInt       int
}

func (m *MeteoData) Average(md *MeteoData) error {

	if m.Station.name != md.Station.name {
		return errors.New("average error: different stations")
	}

	m.TempOut = (m.TempOut + md.TempOut) / 2
	m.HiTemp = (m.HiTemp + md.HiTemp) / 2
	m.LowTemp = (m.LowTemp + md.LowTemp) / 2
	m.OutHum = (m.OutHum + md.OutHum) / 2
	// m.DewPt       =(m.DewPt +md.DewPt)/2
	m.WindSpeed = (m.WindSpeed + md.WindSpeed) / 2
	m.WindRun = (m.WindRun + md.WindRun) / 2
	m.HiSpeed = (m.HiSpeed + md.HiSpeed) / 2
	m.WindChill = (m.WindChill + md.WindChill) / 2
	// m.HeatIndex   =(m.HeatIndex +md.HeatIndex)/2
	// m.THWIndex    =(m.THWIndex +md.THWIndex)/2
	m.Bar = (m.Bar + md.Bar) / 2
	m.Rain = (m.Rain + md.Rain) / 2
	m.RainRate = (m.RainRate + md.RainRate) / 2
	// m.HeatD_D     =(m.HeatD_D +md.HeatD_D)/2
	// m.CoolD_D    =(m.CoolD_D +md.CoolD_D)/2

	return nil
}
