package utils

import (
	"bytes"
	"encoding/json"
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

func LogTestErrorBody(res *http.Response) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	log.Println("--- ERROR RESPONSE ---")
	log.Fatal(string(body))
}
