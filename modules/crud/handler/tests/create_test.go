package crudtests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/exoplanets/modules/crud/handler"
	"github.com/exoplanets/modules/data"
	"github.com/exoplanets/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupPayload(c *gin.Context, payload *data.Exoplanet) {

	b, _ := json.Marshal(&payload)

	c.Request = &http.Request{
		Body:   io.NopCloser(bytes.NewBuffer(b)),
		Header: http.Header{},
		Method: http.MethodPost,
	}

	c.Request.Header.Add("Content-Type", "application/json")

	utils.BindValidationMethods()
}

func TestCreateExoplanet(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	setupPayload(c, &data.Exoplanet{
		Name: "Test Exoplanet", Description: "test description", Distance: 50,
		Radius: 2, Type: "GasGiant",
	})

	handler.CreateExoplanet(c)

	assert.EqualValues(t, http.StatusCreated, w.Code)
	assert.Equal(t, len(data.ExoplanetMapData), 1) // after adding one element we can expect the length will be 1
}

// selecting type terrestrial and not entering the mass so it should give the validation error
func TestCreateExoplanetValidationError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	setupPayload(c, &data.Exoplanet{
		Name: "Test Exoplanet", Description: "test description", Distance: 50,
		Radius: 2, Type: "Terrestrial",
	})

	handler.CreateExoplanet(c)

	assert.EqualValues(t, http.StatusBadRequest, w.Code)
}
