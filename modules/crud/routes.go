package operations

import (
	"github.com/exoplanets/modules/crud/handler"
	"github.com/gin-gonic/gin"
)

// as service name is exoplanet so crud folder by default suggest the exoplanet crud operations
func Routes(router *gin.RouterGroup) {

	router.POST("/", handler.CreateExoplanet)

	router.GET("/list", handler.GetExoplanetList)
	router.GET("/:exoplanetId", handler.GetExoplanetById)

	router.PUT("/:exoplanetId", handler.UpdateExoplanet)

	router.DELETE("/:exoplanetId", handler.DeleteExoplanet)
}
