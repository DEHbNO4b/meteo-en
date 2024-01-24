package meteofile

import (
	"fmt"
	"math"
	"meteo-lightning/internal/domain/models"
	"strconv"
	"time"
)

func meteoToDomain(md meteoData) (models.MeteoData, error) {

	dmd := models.MeteoData{}

	// parsing datetime
	time, err := time.Parse("02.01.06 15:04", md.Date+" "+md.Time)
	if err != nil {
		return dmd, fmt.Errorf("unable to parse date %w", err)

	}

	// parsing temperature
	temp, err := strconv.ParseFloat(md.TempOut, 64)
	if err != nil {
		// return dmd, fmt.Errorf("unable to parse temp %w", err)
		temp = math.NaN()
	}

	// parsing bar
	bar, err := strconv.ParseFloat(md.Bar, 64)
	if err != nil {
		// return dmd, fmt.Errorf("unable to parse bar %w", err)
		bar = math.NaN()
	}
	// parsing precipitation parameters: rain rainRate
	rain, err := strconv.ParseFloat(md.Rain, 64)
	if err != nil {
		// return dmd, fmt.Errorf("unable to parse rain %w", err)
		rain = math.NaN()
	}
	rainRate, err := strconv.ParseFloat(md.RainRate, 64)
	if err != nil {
		// return dmd, fmt.Errorf("unable to parse RainRate %w", err)
		rainRate = math.NaN()
	}

	// parsing wind parameters: speed run chill
	speed, err := strconv.ParseFloat(md.WindSpeed, 64)
	if err != nil {
		// return dmd, fmt.Errorf("unable to parse WindSpeed %w", err)
		speed = math.NaN()
	}
	run, err := strconv.ParseFloat(md.WindRun, 64)
	if err != nil {
		// return dmd, fmt.Errorf("unable to parse WindRun %w", err)
		run = math.NaN()
	}
	chill, err := strconv.ParseFloat(md.WindChill, 64)
	if err != nil {
		// return dmd, fmt.Errorf("unable to parse WindChill %w", err)
		chill = math.NaN()
	}

	dmd.Station = md.FileName
	dmd.Time = time
	dmd.TempOut = temp
	dmd.WindSpeed = speed
	dmd.WindDir = md.WindDir
	dmd.WindRun = run
	dmd.WindChill = chill
	dmd.Bar = bar
	dmd.Rain = rain
	dmd.RainRate = rainRate

	return dmd, nil
}
