package calculations

import (
	"github.com/exoplanets/modules/calculations/handler"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {

	router.POST("/:exoplanetId/fuel/estimation", handler.CalculateExoplanetFuel)

}
