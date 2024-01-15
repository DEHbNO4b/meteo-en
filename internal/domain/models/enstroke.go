package models

import "time"

type StrokeEN struct {
	cloud   bool
	time    time.Time
	nano    int64
	lat     float32
	long    float32
	signal  int16
	height  int16
	sensors int16
	id      int64
}

func (l *StrokeEN) GetCloud() bool {
	return l.cloud
}
func (l *StrokeEN) GetTime() time.Time {
	return l.time
}
func (l *StrokeEN) GetNano() int64 {
	return l.nano
}
func (l *StrokeEN) GetLat() float32 {
	return l.lat
}
func (l *StrokeEN) GetLong() float32 {
	return l.long
}
func (l *StrokeEN) GetSignal() int16 {
	return l.signal
}
func (l *StrokeEN) GetHeight() int16 {
	return l.height
}
func (l *StrokeEN) GetSensors() int16 {
	return l.sensors
}
func (l *StrokeEN) GetID() int64 {
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
func (l *StrokeEN) SetLat(lat float32) {
	l.lat = lat
}
func (l *StrokeEN) SetLong(long float32) {
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
