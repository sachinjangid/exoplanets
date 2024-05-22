package utils

import (
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func BindValidationMethods() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("exoplanetType", ExoplanetType)
		v.RegisterValidation("mass", Mass)
	}
}

func ExoplanetType(fl validator.FieldLevel) bool {
	field := fl.Field().String()

	if strings.TrimSpace(field) == "" {
		return false
	}

	_, valid := AllowedExoplanetType[field]
	return valid
}

func Mass(fl validator.FieldLevel) bool {
	field := fl.Field().Float()

	typeField := fl.Parent().FieldByName("Type").String()

	// if exoplanet is of GasGiant then we can not expect the mass
	if typeField == MassTypeGasGiant {
		return true
	}

	if 0.1 > field || field > 10.0 {
		return false
	}

	return true
}
