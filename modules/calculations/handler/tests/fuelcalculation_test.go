package calculationtests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/exoplanets/modules/calculations/handler"
	"github.com/exoplanets/modules/data"
	"github.com/exoplanets/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupFuelCalculationPayload(c *gin.Context, id string) {

	b, _ := json.Marshal(&handler.FuelCalculationRequest{CrewSize: 2})
	c.Request = &http.Request{
		Body:   io.NopCloser(bytes.NewBuffer(b)),
		Header: http.Header{},
		Method: http.MethodPost,
	}

	c.Params = gin.Params{}
	c.AddParam("exoplanetId", id)

	c.Request.Header.Add("Content-Type", "application/json")

	data.ExoplanetMapData[1] = &data.Exoplanet{
		Name: "Test Exoplanet", Description: "test description", Distance: 50,
		Radius: 1.2, Type: "GasGiant",
	}

	data.ExoplanetMapData[2] = &data.Exoplanet{
		Name: "Test Exoplanet 2", Description: "test description other", Distance: 5,
		Radius: 1.2, Type: "Terrestrial", Mass: 4.0,
	}

	utils.BindValidationMethods()
}

func TestFuelCalculationGasGiants(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	setupFuelCalculationPayload(c, "1")

	handler.CalculateExoplanetFuel(c)

	var got struct {
		Payload handler.FuelCalculateResponse `json:"payload"`
	}

	json.Unmarshal(w.Body.Bytes(), &got)

	assert.EqualValues(t, http.StatusOK, w.Code)
	assert.Equal(t, float64(829.44), got.Payload.FuelCost) // after adding one element we can expect the length will be 1
}

func TestFuelCalculationTerrestrial(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	setupFuelCalculationPayload(c, "2")

	handler.CalculateExoplanetFuel(c)

	var got struct {
		Payload handler.FuelCalculateResponse `json:"payload"`
	}

	json.Unmarshal(w.Body.Bytes(), &got)

	assert.EqualValues(t, http.StatusOK, w.Code)
	assert.Equal(t, float64(1.3), got.Payload.FuelCost) // after adding one element we can expect the length will be 1
}
