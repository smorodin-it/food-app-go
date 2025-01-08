package utils

import (
	"fmt"
	"food-backend/src/constants"
	"food-backend/src/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GetAuthTokenDuration() time.Time {
	return time.Now().Add(constants.AuthTokenDuration)
}

func GetAuthRefreshTokenDuration() time.Time {
	return time.Now().Add(constants.AuthRefreshTokenDuration)
}

func SetTokensToResponse(ctx *fiber.Ctx, tokens responses.ResponseTokens) {
	// Set refresh token to server cookie
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.AuthRefreshTokenField,
		Value:    tokens.RefreshToken,
		Path:     "/",
		Domain:   ctx.Hostname(),
		Expires:  GetAuthRefreshTokenDuration(),
		HTTPOnly: true,
		SameSite: "strict",
	})

	// Set access token to Authorization header
	ctx.Set("Authorization", fmt.Sprintf("Bearer %s", tokens.Token))
}

func CalcPagination(page int, perPage int) (limit int, offset int, withPagination bool) {
	if page >= 1 && perPage > 0 {
		offset = perPage * (page - 1)
		withPagination = true
	} else {
		offset = 0
		withPagination = false
	}

	limit = perPage

	return limit, offset, withPagination
}

// TODO: Check what to use this function or method from services.AuthService
func GetUserIDFromToken(ctx *fiber.Ctx) string {
	return ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["id"].(string)
}

func GetResponseData(ctx *fiber.Ctx, data any) error {
	return ctx.Status(fiber.StatusOK).JSON(&responses.ResponseApi[any]{
		Data:  data,
		Error: nil,
	})
}

func GetResponseError(ctx *fiber.Ctx, status int, err error) error {
	e := err.Error()

	return ctx.Status(status).JSON(&responses.ResponseApi[any]{
		Data:  nil,
		Error: &e,
	})
}

func GetResponseAdd(ctx *fiber.Ctx, id string) error {
	return GetResponseData(ctx, responses.ResponseAdd{
		Id: id,
	})
}

func GetResponseStatus(ctx *fiber.Ctx, status bool) error {
	return GetResponseData(ctx, responses.ResponseStatus{
		Status: status,
	})
}
