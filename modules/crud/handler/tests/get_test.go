package crudtests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/exoplanets/modules/crud/handler"
	"github.com/exoplanets/modules/data"
	"github.com/exoplanets/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupGetPayload(c *gin.Context, id string) {

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

	data.ExoplanetMapData[2] = &data.Exoplanet{
		Name: "Test Exoplanet 2", Description: "test description other", Distance: 5,
		Radius: 1.2, Type: "Terrestrial", Mass: 4.0,
	}

	utils.BindValidationMethods()
}

func TestGetExoplanet(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	setupGetPayload(c, "1")

	handler.GetExoplanetById(c)

	expectedData := data.ExoplanetMapData[1]
	var got struct {
		Payload data.Exoplanet `json:"payload"`
	}

	json.Unmarshal(w.Body.Bytes(), &got)

	assert.EqualValues(t, http.StatusOK, w.Code)
	assert.EqualValues(t, expectedData.Name, got.Payload.Name) // compare the data coming from the response

}

func TestGetListExoplanet(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	setupGetPayload(c, "1")

	handler.GetExoplanetList(c)

	var got struct {
		Payload data.Exoplanets `json:"payload"`
	}

	json.Unmarshal(w.Body.Bytes(), &got)

	assert.EqualValues(t, http.StatusOK, w.Code)
	assert.EqualValues(t, len(got.Payload), 2) // we set two exoplanets we can expect 2

}
