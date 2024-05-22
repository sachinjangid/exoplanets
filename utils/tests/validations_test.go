package utiltests

import (
	"fmt"
	"testing"

	"github.com/exoplanets/utils"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestCustomValidationsMass(t *testing.T) {

	validate := validator.New()

	err := validate.RegisterValidation("mass", utils.Mass)
	if err != nil {
		fmt.Println(err)
	}

	testInputs := []struct {
		Name     string
		Mass     float64 `json:"mass,omitempty" validate:"mass"`
		Type     string
		Expected bool
	}{
		{Name: "Violating min limit", Type: utils.MassTypeTerrestrial, Mass: 0.0, Expected: false},
		{Name: "Passing min limit", Type: utils.MassTypeGasGiant, Mass: 0.0, Expected: true},
		{Name: "Should pass value 2.0", Type: utils.MassTypeTerrestrial, Mass: 2.0, Expected: true},
		{Name: "Should pass value 3.1", Type: utils.MassTypeTerrestrial, Mass: 3.1, Expected: true},
		{Name: "Violating max limit", Type: utils.MassTypeTerrestrial, Mass: 11.1, Expected: false},
	}

	validate.Struct(testInputs[0])
	for _, input := range testInputs {
		t.Run(input.Name, func(t *testing.T) {
			err := validate.Struct(input)
			if input.Expected == true {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}

}

func TestCustomValidationsType(t *testing.T) {

	validate := validator.New()

	err := validate.RegisterValidation("exoplanetType", utils.ExoplanetType)
	if err != nil {
		fmt.Println(err)
	}

	testInputs := []struct {
		Name    string
		Type    string `validate:"exoplanetType"`
		IsValid bool
	}{
		{Name: "Pass for GasGiant", Type: utils.MassTypeGasGiant, IsValid: true},
		{Name: "Pass for Terrestrial", Type: utils.MassTypeTerrestrial, IsValid: true},
		{Name: "Fail for TestPlanet", Type: "TestPlanet", IsValid: false},
	}

	validate.Struct(testInputs[0])
	for _, input := range testInputs {
		t.Run(input.Name, func(t *testing.T) {
			err := validate.Struct(input)
			if input.IsValid == true {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}

}
