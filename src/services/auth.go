package services

import (
	"food-backend/src/constants"
	"food-backend/src/domains"
	"food-backend/src/responses"
	"food-backend/src/utils"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type AuthService struct {
}

func (r AuthService) GenerateTokens(user *domains.User) (*responses.ResponseTokens, error) {
	// Generate access token
	token := jwt.New(jwt.SigningMethodHS512)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
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
