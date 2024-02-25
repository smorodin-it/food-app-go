package handlers

import (
	"food-backend/src/forms"
	"food-backend/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"time"
)

var testUserCredentials = &forms.FormAuth{
	Email:    "test@test.loc",
	Password: "12332145",
}

func TestAuthUserHandler(t *testing.T) {
	app := SetupTestApp()

	res := utils.MakeTestRequest(t, app, "POST", "/api/auth", testUserCredentials)

	// Verify status code
	pass := assert.Equal(t, fiber.StatusOK, res.StatusCode, "should have 200 status code")
	if !pass {
		return
	}

	//	Verify headers
	pass = assert.Equal(t, true, len(res.Header.Get("Authorization")) > 0, "should have authorization token in header")
	if !pass {
		return
	}

	cookies := res.Cookies()
	hasRefreshToken := false
	for _, cookie := range cookies {
		if cookie.Name == "refreshToken" {
			hasRefreshToken = true
			break
		}
	}
	assert.Equal(t, true, hasRefreshToken, "should have refresh token in cookies")
}

func TestRefreshTokensHandler(t *testing.T) {
	app := SetupTestApp()

	// Authorize to get tokens
	res := utils.MakeTestRequest(t, app, "POST", "/api/auth", testUserCredentials)
	accessToken := res.Header.Get("Authorization")
	cookies := res.Cookies()
	var refreshTokenCookie *http.Cookie = nil
	for _, cookie := range cookies {
		if cookie.Name == "refreshToken" {
			refreshTokenCookie = cookie
			break
		}
	}

	// Without sleep, tokens are same
	time.Sleep(time.Millisecond * 500)

	// Refresh tokens
	req, _ := http.NewRequest("POST", "/api/auth/refresh", nil)
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(refreshTokenCookie)

	res, err := app.Test(req, -1)
	pass := assert.Equal(t, false, err != nil, "app.Test should not return error")
	if !pass {
		log.Fatal(err.Error())
	}

	// Verify status code
	pass = assert.Equal(t, fiber.StatusOK, res.StatusCode, "should have 200 status code")
	if !pass {
		utils.LogTestErrorBody(res)
	}

	//	Verify headers
	newAccessToken := res.Header.Get("Authorization")
	pass = assert.Equal(t, true, len(newAccessToken) > 0, "should have authorization token in header")
	if !pass {
		return
	}

	pass = assert.NotEqual(t, accessToken, newAccessToken, "access token should not be same as from auth request")
	if !pass {
		log.Println("Access Token: ", accessToken)
		log.Println("New Access Token: ", newAccessToken)
		return
	}

	cookies = res.Cookies()
	newRefreshToken := ""
	for _, cookie := range cookies {
		if cookie.Name == "refreshToken" {
			newRefreshToken = cookie.Value
			break
		}
	}

	pass = assert.Equal(t, true, len(newRefreshToken) > 0, "should have refresh token in cookies")
	if !pass {
		return
	}

	assert.NotEqual(t, refreshTokenCookie.Value, newRefreshToken, "refresh token should not be same as from auth request")
}
