package services

import (
	"errors"
	"food-backend/src/constants"
	"food-backend/src/domains"
	"food-backend/src/forms"
	"food-backend/src/responses"
	"food-backend/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService struct {
}

func (s AuthService) GenerateTokens(user *domains.User) (*responses.ResponseTokens, error) {
	// Generate access token
	token := jwt.New(jwt.SigningMethodHS512)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.UserID
	claims["email"] = user.Email
	claims["exp"] = utils.GetAuthTokenDuration().Unix()

	at, err := token.SignedString([]byte(constants.AuthSignKey))
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	rt := uuid.New().String()

	return &responses.ResponseTokens{Token: at, RefreshToken: rt}, nil
}

func (s AuthService) GetUserIdFromToken(ctx *fiber.Ctx) (id interface{}, err error) {
	f := new(forms.AccessTokenForm)

	err = ctx.CookieParser(f)
	if err != nil {
		return nil, err
	}

	t, err := jwt.Parse(f.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.AuthSignKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok {
		return claims["id"], nil
	} else {
		return nil, errors.New("can't parse token")
	}
}
