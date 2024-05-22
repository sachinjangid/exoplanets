package handler

import (
	"net/http"
	"strconv"

	"github.com/exoplanets/modules/data"
	"github.com/exoplanets/utils"
	"github.com/gin-gonic/gin"
)

func GetExoplanetList(ctx *gin.Context) {

	result := data.Exoplanets{}
	for _, exoplanet := range data.ExoplanetMapData {
		result = append(result, exoplanet)
	}

	utils.ReturnResponse(ctx, http.StatusOK, utils.ExoplanetListed, result)
}

func GetExoplanetById(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Params.ByName("exoplanetId"))
	if err != nil { // indicates some wrong type of data is been passed
		utils.ReturnResponse(ctx, http.StatusBadRequest, utils.InvalidExoplanetId, nil)
		return
	}

	exoplanetData, ok := data.ExoplanetMapData[id]
	if !ok {
		utils.ReturnResponse(ctx, http.StatusBadRequest, utils.InvalidExoplanetId, nil)
		return
	}

	utils.ReturnResponse(ctx, http.StatusOK, utils.ExoplanetListed, exoplanetData)
}
