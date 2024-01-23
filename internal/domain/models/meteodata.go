package models

import "time"

type MeteoData struct {
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
