package utils

var ServiceName string = "exoplanet"

const MassTypeGasGiant string = "GasGiant"
const MassTypeTerrestrial string = "Terrestrial"

var AllowedExoplanetType = map[string]struct{}{
	MassTypeGasGiant:    {},
	MassTypeTerrestrial: {},
}

// error messages
var ValidationError string = "Validation Error"
var InvalidExoplanetPayload string = "Invalid request data"
var InvalidExoplanetId string = "Invalid exoplanet id"

// success messages
var ExoplanetCreated string = "Exoplanet Stored Successfully!!"
var ExoplanetListed string = "Exoplanet Fetched Successfully!!"
var ExoplanetDeleted string = "Exoplanet Deleted Successfully!!"
var ExoplanetUpdated string = "Exoplanet Updated Successfully!!"

var CalculationDone string = "Calculation Done!!"
