package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/exoplanets/modules/data"
	"github.com/exoplanets/utils"
	"github.com/gin-gonic/gin"
)

func UpdateExoplanet(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Params.ByName("exoplanetId"))
	if err != nil { // indicates some wrong type of data is been passed
		utils.ReturnResponse(ctx, http.StatusBadRequest, utils.InvalidExoplanetId, nil)
		return
	}

	var request data.Exoplanet

	if err := ctx.Bind(&request); err != nil {
		utils.HandleValidationErr(err, ctx)
		return
	}

	prev, ok := data.ExoplanetMapData[id]
	if !ok {
		utils.ReturnResponse(ctx, http.StatusBadRequest, utils.InvalidExoplanetId, nil)
		return
	}

	// overrite the values that we are not expecting in the request
	currentTime := time.Now()
	request.Id = id
	request.FoundAt = prev.FoundAt
	request.UpdatedAt = &currentTime

	data.ExoplanetMapData[id] = &request

	utils.ReturnResponse(ctx, http.StatusOK, utils.ExoplanetUpdated, nil)
}
