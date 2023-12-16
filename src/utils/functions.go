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
