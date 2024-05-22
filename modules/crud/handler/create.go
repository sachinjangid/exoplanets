package handler

import (
	"net/http"
	"time"

	"github.com/exoplanets/modules/data"
	"github.com/exoplanets/utils"
	"github.com/gin-gonic/gin"
)

/**
I generally follow the practice to create any module in
handler, service, repository folders at minimum
but here logics are straight forward so did not made it complicated unnecessary
also here I am using function as a handler
but in general I prefer to use the methods first by creating the struct ex:
for this create operation we could have made a struct

type CreateExoplanetHandler struct {}

and then could define the below function as follows

func (c *CreateExoplanetHandler) CreateExoplanet(ctx *gin.Context) {}

but this practice I follow when we really need to import different packages like service, repository
or other helper

I have tried to make the code more simpler
*/

func CreateExoplanet(ctx *gin.Context) {
	var request data.Exoplanet

	if err := ctx.Bind(&request); err != nil {
		utils.HandleValidationErr(err, ctx)
		return
	}

	currentTime := time.Now()
	exoplanetId := len(data.ExoplanetMapData) + 1
	request.Id = exoplanetId
	request.FoundAt = &currentTime
	request.UpdatedAt = &currentTime

	data.ExoplanetMapData[exoplanetId] = &request

	utils.ReturnResponse(ctx, http.StatusCreated, utils.ExoplanetCreated, nil)
}
