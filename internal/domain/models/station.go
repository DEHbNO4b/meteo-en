package models

type Station struct {
	id        int64
	name      string
	latitude  float64
	longitude float64
}

func (s *Station) SetID(id int64) {
	s.id = id
}
func (s *Station) SetName(n string) {
	s.name = n
}
func (s *Station) SetLat(l float64) {
	s.latitude = l
}
func (s *Station) SetLong(l float64) {
	s.longitude = l
}
func (s *Station) ID() int64 {
	return s.id
}
func (s *Station) Name() string {
	return s.name
}
func (s *Station) Lat() float64 {
	return s.latitude
}
func (s *Station) Long() float64 {
	return s.longitude
}
