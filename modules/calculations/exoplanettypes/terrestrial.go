package exoplanettypes

import "math"

type Terrestrial struct {
	Mass     float64
	Distance int
	Radiaus  float64
}

func (t *Terrestrial) CalculateFuel() float64 {
	gravity := (t.Mass / math.Pow(t.Radiaus, 2)) // gravity

	fuelUnit := (float64(t.Distance) / math.Pow(gravity, 2))

	return math.Round(fuelUnit*100) / 100
}
