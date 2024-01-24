package models

import "time"

type StrokeEN struct {
	cloud   bool
	time    time.Time
	nano    int64
	lat     float64
	long    float64
	signal  int16
	height  int16
	sensors int16
	id      int64
}

func (l *StrokeEN) Cloud() bool {
	return l.cloud
}
func (l *StrokeEN) Time() time.Time {
	return l.time
}
func (l *StrokeEN) Nano() int64 {
	return l.nano
}
func (l *StrokeEN) Lat() float64 {
	return l.lat
}
func (l *StrokeEN) Long() float64 {
	return l.long
}
func (l *StrokeEN) Signal() int16 {
	return l.signal
}
func (l *StrokeEN) Height() int16 {
	return l.height
}
func (l *StrokeEN) Sensors() int16 {
	return l.sensors
}
func (l *StrokeEN) ID() int64 {
	return l.id
}

func (l *StrokeEN) SetCloud(c bool) {
	l.cloud = c
}
func (l *StrokeEN) SetTime(t time.Time) {
	l.time = t
}
func (l *StrokeEN) SetNano(n int64) {
	l.nano = n
}
func (l *StrokeEN) SetLat(lat float64) {
	l.lat = lat
}
func (l *StrokeEN) SetLong(long float64) {
	l.long = long
}
func (l *StrokeEN) SetSignal(s int16) {
	l.signal = s
}
func (l *StrokeEN) SetHeight(h int16) {
	l.height = h
}
func (l *StrokeEN) SetSensors(s int16) {
	l.sensors = s
}
func (l *StrokeEN) SetID(id int64) {
	l.id = id
}
