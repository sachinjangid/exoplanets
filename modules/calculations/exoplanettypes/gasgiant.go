package exoplanettypes

import "math"

type GasGiant struct {
	Distance int
	Radiaus  float64
}

func (g *GasGiant) CalculateFuel() float64 {
	gravity := (0.5 / math.Pow(g.Radiaus, 2)) // gravity

	fuelUnit := (float64(g.Distance) / math.Pow(gravity, 2))

	return math.Round(fuelUnit*100) / 100
}
