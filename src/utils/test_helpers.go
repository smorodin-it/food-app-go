package utils

import (
	"bytes"
	"encoding/json"
	"food-backend/src/forms"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"testing"
)

func MakeTestRequest(t *testing.T, app *fiber.App, method string, url string, body any) (res *http.Response) {
	readBody, err := json.Marshal(body)
	assert.Equal(t, true, err == nil, "should marshal body")

	jsonVal := bytes.NewBuffer(readBody)
	req, _ := http.NewRequest(method, url, jsonVal)
	req.Header.Set("Content-Type", "application/json")

	res, err = app.Test(req, -1)
	pass := assert.Equal(t, false, err != nil, "app.Test should not return error")
	if !pass {
		return nil
	}

	return res
}

func MakeAuthTestRequest(t *testing.T, app *fiber.App, method string, url string, body any) (res *http.Response) {
	// Auth user
	readBody, err := json.Marshal(&forms.FormAuth{
		Email:    "test@test.loc",
		Password: "12332145",
	})
	jsonVal := bytes.NewBuffer(readBody)
	req, err := http.NewRequest("POST", "/api/auth", jsonVal)
	if err != nil {
		panic(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	res, err = app.Test(req, -1)

	// Get auth header
	authHeader := res.Header.Get("Authorization")

	// Make API request
	readBody, err = json.Marshal(body)
	assert.Equal(t, true, err == nil, "should marshal body")

	jsonVal = bytes.NewBuffer(readBody)
	req, err = http.NewRequest(method, url, jsonVal)
	if err != nil {
		panic(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authHeader)

	res, err = app.Test(req, -1)
	pass := assert.Equal(t, false, err != nil, "app.Test should not return error")
	if !pass {
		panic(err.Error())
	}

	return res
}

func LogTestErrorBody(res *http.Response) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	log.Println("--- ERROR RESPONSE BODY ---")
	log.Fatal(string(body))
}
