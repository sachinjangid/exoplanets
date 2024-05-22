package typestests

import (
	"testing"

	"github.com/exoplanets/modules/calculations/exoplanettypes"
	"github.com/stretchr/testify/assert"
)

func TestTerrestrialFuelCostUnit(t *testing.T) {
	mockExoplanet := exoplanettypes.Terrestrial{
		Distance: 15,
		Radiaus:  4.6,
		Mass:     3.1,
	}

	fuelCostUnit := mockExoplanet.CalculateFuel()
	assert.Equal(t, 698.87, fuelCostUnit)

}
