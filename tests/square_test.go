package tests

import (
	"fmt"
	"math"
	"meteo-lightning/internal/domain/models"
	"testing"
)

func TestCalculateSquareCorners(t *testing.T) {

	station := models.NewStation(40, 44)

	square := station.CalculateSquareCorners(10)

	fmt.Printf("%+v\n", square)

	want := 20.0

	got := square.UpperLeft.DistanceTo(square.UpperRight)

	if math.Abs(got-want) > 0.0 {
		t.Errorf("got %v; want %v", got, want)
	}

}
