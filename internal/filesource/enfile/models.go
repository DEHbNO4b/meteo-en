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
	Signal  int
	Height  int
	Sensors int16
}
