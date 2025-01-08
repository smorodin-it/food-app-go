package services

import (
	"errors"
	"food-backend/src/constants"
	"food-backend/src/domains"
	"food-backend/src/forms"
	"food-backend/src/repositories"
	"food-backend/src/responses"
	"food-backend/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type AuthService interface {
	GetUserIdFromToken(ctx *fiber.Ctx) (id interface{}, err error)
	AuthUser(form forms.AuthForm) (tokens *responses.ResponseTokens, err error)
	RefreshTokens(form forms.RefreshForm) (tokens *responses.ResponseTokens, err error)
}

type authService struct {
	r repositories.UserRepository
}

func (s authService) generateTokens(user *domains.User) (*responses.ResponseTokens, error) {
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

// TODO: Check what to use this method or function from utils.GetUserIDFromToken
func (s authService) GetUserIdFromToken(ctx *fiber.Ctx) (id interface{}, err error) {
	token := ctx.Get("Authorization", "")
	token = strings.Split(token, "Bearer ")[1]
	if token == "" {
		return nil, errors.New("token is empty")
	}

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
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

func (s authService) AuthUser(form forms.AuthForm) (tokens *responses.ResponseTokens, err error) {
	//	Get user by email
	user, err := s.r.GetByEmail(form.Email)
	if err != nil {
		return nil, err
	}

	// Check user password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(form.Password))
	if err != nil {
		return nil, err
	}

	// Generate tokens
	tokens, err = s.generateTokens(user)
	if err != nil {
		return nil, err
	}

	err = s.r.SetRefreshToken(user.UserID, tokens.RefreshToken)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (s authService) RefreshTokens(form forms.RefreshForm) (tokens *responses.ResponseTokens, err error) {
	// Get user from db by refresh token
	user, err := s.r.GetByRefreshToken(form.RefreshToken)
	if err != nil {
		return nil, err
	}

	// Generate and update tokens
	tokens, err = s.generateTokens(user)
	if err != nil {
		return nil, err
	}

	// Set new refresh token to user
	err = s.r.SetRefreshToken(user.UserID, tokens.RefreshToken)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func NewAuthService(repository repositories.UserRepository) AuthService {
	return &authService{
		r: repository,
	}
}
