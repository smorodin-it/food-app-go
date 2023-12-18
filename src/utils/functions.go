package utils

import (
	"food-backend/src/constants"
	"food-backend/src/responses"
	"github.com/gofiber/fiber/v2"
	"time"
)

func GetAuthTokenDuration() time.Time {
	return time.Now().Add(constants.AuthTokenDuration)
}

func GetAuthRefreshTokenDuration() time.Time {
	return time.Now().Add(constants.AuthRefreshTokenDuration)
}

func SetTokensToCookies(ctx *fiber.Ctx, t responses.ResponseTokens) {
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.AuthRefreshTokenField,
		Value:    t.RefreshToken,
		Path:     "/",
		Domain:   ctx.Hostname(),
		Expires:  GetAuthRefreshTokenDuration(),
		HTTPOnly: true,
		SameSite: "strict",
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.AuthAccessTokenField,
		Value:    t.Token,
		Path:     "/",
		Domain:   ctx.Hostname(),
		Expires:  GetAuthTokenDuration(),
		HTTPOnly: true,
		SameSite: "strict",
	})
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
