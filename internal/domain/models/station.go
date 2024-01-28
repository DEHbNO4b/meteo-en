package models

import "math"

type Station struct {
	id   int64
	name string
	Point
}

func NewStation(lat, long float64) *Station {
	return &Station{
		Point: Point{lat: lat, long: long},
	}
}

func (s *Station) SetID(id int64) {
	s.id = id
}
func (s *Station) SetName(n string) {
	s.name = n
}

func (s *Station) ID() int64 {
	return s.id
}
func (s *Station) Name() string {
	return s.name
}

// Функция для вычисления координат углов квадрата
func (s *Station) CalculateSquareCorners(radius float64) Square {
	// Вычисляем смещение в градусах для 100 км
	offsetLat := radius / 111.0
	offsetLong := radius / (111.0 * math.Cos(s.Lat()*(math.Pi/180.0)))

	// Вычисляем координаты углов квадрата
	upperLeft := Point{}
	upperLeft.SetLat(s.Lat() + offsetLat)
	upperLeft.SetLong(s.Long() - offsetLong)

	upperRight := Point{}
	upperRight.SetLat(s.Lat() + offsetLat)
	upperRight.SetLong(s.Long() + offsetLong)

	lowerLeft := Point{}
	lowerLeft.SetLat(s.Lat() - offsetLat)
	lowerLeft.SetLong(s.Long() - offsetLong)

	lowerRight := Point{}
	lowerRight.SetLat(s.Lat() - offsetLat)
	lowerRight.SetLong(s.Long() + offsetLong)

	// Формируем структуру Square с координатами углов
	square := Square{
		UpperLeft:  upperLeft,
		UpperRight: upperRight,
		LowerLeft:  lowerLeft,
		LowerRight: lowerRight,
	}

	return square
}
