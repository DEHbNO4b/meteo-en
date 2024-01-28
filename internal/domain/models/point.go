package models

import (
	"github.com/umahmood/haversine"
)

type Point struct {
	lat  float64
	long float64
}

func (s *Point) Lat() float64 {
	return s.lat
}
func (s *Point) SetLat(l float64) {
	s.lat = l
}
func (s *Point) Long() float64 {
	return s.long
}
func (s *Point) SetLong(l float64) {
	s.long = l
}
func (s *Point) DistanceTo(other Point) float64 {
	this := haversine.Coord{Lat: s.lat, Lon: s.long}
	oth := haversine.Coord{Lat: other.Lat(), Lon: other.Long()}
	_, km := haversine.Distance(this, oth)
	return km
}
