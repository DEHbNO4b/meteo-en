package enfile

import (
	"time"
)

type stroke struct {
	Cloud   bool
	Time    time.Time
	Nano    int64
	Lat     float64
	Long    float64
	Signal  int16
	Height  int16
	Sensors int16
}
