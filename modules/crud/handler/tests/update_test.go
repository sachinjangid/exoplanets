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

func setupUpdatePayload(c *gin.Context, payload *data.Exoplanet, id string) {

	b, _ := json.Marshal(&payload)

	c.Request = &http.Request{
		Body:   io.NopCloser(bytes.NewBuffer(b)),
		Header: http.Header{},
		Method: http.MethodPost,
	}

	c.Params = gin.Params{}
	c.AddParam("exoplanetId", id)

	c.Request.Header.Add("Content-Type", "application/json")

	utils.BindValidationMethods()
}

func TestUpdateExoplanet(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// set default data so that we can update it later
	data.ExoplanetMapData[1] = &data.Exoplanet{Name: "Test Exoplanet", Description: "test description",
		Distance: 50, Radius: 2, Type: "GasGiant",
	}

	newPayload := data.Exoplanet{Name: "Test Exoplanet Updated", Description: "test description updated",
		Distance: 20, Radius: 1.2, Type: "GasGiant",
	}
	setupUpdatePayload(c, &newPayload, "1")

	handler.UpdateExoplanet(c)

	updatedExoplanet := data.ExoplanetMapData[1]

	assert.EqualValues(t, http.StatusOK, w.Code)
	assert.Equal(t, newPayload.Name, updatedExoplanet.Name)
	assert.Equal(t, newPayload.Distance, updatedExoplanet.Distance)
	assert.Equal(t, newPayload.Description, updatedExoplanet.Description)
	assert.Equal(t, newPayload.Radius, updatedExoplanet.Radius)
}
