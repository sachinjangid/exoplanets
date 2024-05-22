package main

import (
	"net/http"

	"github.com/exoplanets/modules/calculations"
	exoplanetCRUD "github.com/exoplanets/modules/crud"
	"github.com/exoplanets/utils"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	routeGroup := router.Group(utils.ServiceName)

	utils.BindValidationMethods()

	exoplanetCRUD.Routes(routeGroup)
	calculations.Routes(routeGroup)

	http.ListenAndServe(":8080", router)
}
