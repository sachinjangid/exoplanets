package crudtests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/exoplanets/modules/crud/handler"
	"github.com/exoplanets/modules/data"
	"github.com/exoplanets/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupDeletePayload(c *gin.Context, id string) {

	c.Request = &http.Request{
		Header: http.Header{},
		Method: http.MethodPost,
	}

	c.Params = gin.Params{}
	c.AddParam("exoplanetId", id)

	c.Request.Header.Add("Content-Type", "application/json")

	data.ExoplanetMapData[1] = &data.Exoplanet{
		Name: "Test Exoplanet", Description: "test description", Distance: 50,
		Radius: 2, Type: "GasGiant",
	}

	utils.BindValidationMethods()
}

func TestDeleteExoplanet(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	setupDeletePayload(c, "1")

	assert.Equal(t, len(data.ExoplanetMapData), 1) // before deletion the size of overall data was 1
	handler.DeleteExoplanet(c)

	assert.EqualValues(t, http.StatusOK, w.Code)
	assert.Equal(t, len(data.ExoplanetMapData), 0) // after deletion it has become 0
}
