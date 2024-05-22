package handler

import (
	"net/http"
	"strconv"

	"github.com/exoplanets/modules/calculations/exoplanettypes"
	"github.com/exoplanets/modules/data"
	"github.com/exoplanets/utils"
	"github.com/gin-gonic/gin"
)

func CalculateExoplanetFuel(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Params.ByName("exoplanetId"))
	if err != nil { // indicates some wrong type of data is been passed
		utils.ReturnResponse(ctx, http.StatusBadRequest, utils.InvalidExoplanetId, nil)
		return
	}

	var request FuelCalculationRequest
	if err := ctx.Bind(&request); err != nil {
		utils.HandleValidationErr(err, ctx)
		return
	}

	exoplanetData, ok := data.ExoplanetMapData[id]
	if !ok {
		utils.ReturnResponse(ctx, http.StatusBadRequest, utils.InvalidExoplanetId, nil)
		return
	}

	calculator := exoplanettypes.FindExoplanetCalculator(exoplanetData)
	fuelUnit := calculator.CalculateFuel()

	response := FuelCalculateResponse{
		FuelCost: fuelUnit * float64(request.CrewSize),
	}

	utils.ReturnResponse(ctx, http.StatusOK, utils.CalculationDone, &response)
}
