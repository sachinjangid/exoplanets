package handler

import (
	"net/http"
	"strconv"

	"github.com/exoplanets/modules/data"
	"github.com/exoplanets/utils"
	"github.com/gin-gonic/gin"
)

func DeleteExoplanet(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Params.ByName("exoplanetId"))
	if err != nil { // indicates some wrong type of data is been passed
		utils.ReturnResponse(ctx, http.StatusBadRequest, utils.InvalidExoplanetId, nil)
		return
	}

	if _, ok := data.ExoplanetMapData[id]; !ok {
		utils.ReturnResponse(ctx, http.StatusBadRequest, utils.InvalidExoplanetId, nil)
		return
	}

	delete(data.ExoplanetMapData, id)

	utils.ReturnResponse(ctx, http.StatusOK, utils.ExoplanetDeleted, nil)
}
