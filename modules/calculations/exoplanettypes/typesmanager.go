package exoplanettypes

import (
	"github.com/exoplanets/modules/data"
	"github.com/exoplanets/utils"
)

type Calculator interface {
	CalculateFuel() float64
}

func FindExoplanetCalculator(exoplanet *data.Exoplanet) Calculator {
	if exoplanet.Type == utils.MassTypeGasGiant {
		return &GasGiant{Distance: exoplanet.Distance, Radiaus: exoplanet.Radius}
	}

	return &Terrestrial{Distance: exoplanet.Distance, Radiaus: exoplanet.Radius, Mass: exoplanet.Mass}
}
