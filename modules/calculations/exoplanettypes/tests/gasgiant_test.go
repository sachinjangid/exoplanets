package typestests

import (
	"testing"

	"github.com/exoplanets/modules/calculations/exoplanettypes"
	"github.com/stretchr/testify/assert"
)

func TestGasGiantFuelCostUnit(t *testing.T) {
	mockExoplanet := exoplanettypes.GasGiant{
		Distance: 15,
		Radiaus:  4.6,
	}

	fuelCostUnit := mockExoplanet.CalculateFuel()
	assert.Equal(t, 26864.74, fuelCostUnit)

}
