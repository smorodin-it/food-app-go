package handlers

import (
	"encoding/json"
	"food-backend/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"reflect"
	"testing"
)

func TestIngredientList(t *testing.T) {
	app := SetupTestApp()

	res := utils.MakeAuthTestRequest(t, app, "GET", "/api/ingredient", nil)

	// Verify status code
	pass := assert.Equal(t, fiber.StatusOK, res.StatusCode, "should have 200 status code")
	if !pass {
		utils.LogTestErrorBody(res)
	}

	expectedResBody := &[]struct {
		Id           string  `json:"id"`
		Name         string  `json:"name"`
		Manufacturer string  `json:"manufacturer"`
		Barcode      *string `json:"barcode"`
		Proteins     int     `json:"proteins"`
		Carbs        int     `json:"carbs"`
		Fats         int     `json:"fats"`
		Calories     int     `json:"calories"`
	}{}

	body, err := io.ReadAll(res.Body)
	pass = assert.Equal(t, true, err == nil, "should read body")
	if !pass {
		utils.LogTestErrorBody(res)
	}

	err = json.Unmarshal(body, expectedResBody)

	for _, ingredient := range *expectedResBody {
		assert.NotNil(t, ingredient.Id, "should have Id field")
		assert.NotNil(t, ingredient.Name, "should have Name field")
		assert.NotNil(t, ingredient.Manufacturer, "should have Manufacturer field")
		assert.Equal(t, true, reflect.ValueOf(ingredient.Barcode).Type().String() == "*string", "should have Barcode field")
		assert.NotNil(t, ingredient.Proteins, "should have Proteins field")
		assert.NotNil(t, ingredient.Carbs, "should have Carbs field")
		assert.NotNil(t, ingredient.Fats, "should have Fats field")
		assert.NotNil(t, ingredient.Calories, "should have Calories field")
	}

	pass = assert.Equal(t, true, err == nil, "should unmarshal response")
	if !pass {
		utils.LogTestErrorBody(res)
	}
}

func TestIngredientCreate(t *testing.T) {
}

func TestIngredientRetrieve(t *testing.T) {
}

func TestIngredientUpdate(t *testing.T) {
}
