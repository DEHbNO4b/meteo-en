package models

import "time"

type MeteoData struct {
	Date         time.Time
	Time         time.Time
	TempOut      float32
	HiTemp       float32
	LowTemp      float32
	OutHum       int
	DewPt        float32
	WindSpeed    float32
	WindDir      string
	WindRun      float32
	HiSpeed      float32
	HiDir        string
	WindChill    float32
	HeatIndex    float32
	THWIndex     float32
	Bar          float32
	Rain         float32
	RainRate     float32
	HeatD_D      float32
	CoolD_D      float32
	InTemp       float32
	InHum        float32
	InDew        float32
	InHeat       float32
	InEMC        float32
	InAirDensity float32
	WindSamp     float32
	WindTx       float32
	ISSRecept    float32
	ArcInt       int
}
