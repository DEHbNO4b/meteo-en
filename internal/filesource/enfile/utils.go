package enfile

import (
	"fmt"
	"meteo-lightning/internal/domain/models"
	"strconv"
	"time"
)

func enToDomain(s stroke) models.StrokeEN {

	stroke := models.StrokeEN{}

	stroke.SetCloud(s.Cloud)
	stroke.SetTime(s.Time)
	stroke.SetNano(s.Nano)
	stroke.SetLat(s.Lat)
	stroke.SetSignal(s.Signal)
	stroke.SetHeight(s.Height)
	stroke.SetSensors(s.Sensors)

	return stroke
}

func makeStroke(rec []string) (stroke, error) {

	s := stroke{}

	if len(rec[0]) == 13 {
		s.Cloud = true
	}

	time, err := time.Parse("2006-01-02 15:04:05", rec[1])
	if err != nil {
		return s, fmt.Errorf("unable to parse time %w", err)

	}

	nano, err := strconv.Atoi(rec[2])
	if err != nil {
		return s, fmt.Errorf("unable to parse nano %w", err)
	}
	lat, err := strconv.ParseFloat(rec[3], 64)
	if err != nil {
		return s, fmt.Errorf("unable to parse latitude %w", err)
	}
	long, err := strconv.ParseFloat(rec[4], 64)
	if err != nil {
		return s, fmt.Errorf("unable to parse longitude %w", err)
	}
	signal, err := strconv.Atoi(rec[5])
	if err != nil {
		return s, fmt.Errorf("unable to parse signal %w", err)
	}
	height, err := strconv.Atoi(rec[6])
	if err != nil {
		return s, fmt.Errorf("unable to parse signal %w", err)
	}
	sensors, err := strconv.Atoi(rec[7])
	if err != nil {
		return s, fmt.Errorf("unable to parse signal %w", err)
	}

	s.Time = time
	s.Nano = int64(nano)
	s.Lat = lat
	s.Long = long
	s.Signal = int16(signal)
	s.Height = int16(height)
	s.Sensors = int16(sensors)

	return s, nil

}
