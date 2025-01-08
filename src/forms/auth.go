package forms

import "errors"

type AuthForm struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (f AuthForm) Validate() error {
	return nil
}

type RefreshForm struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

func (f RefreshForm) Validate() error {
	if f.RefreshToken == "" {
		return errors.New("bad refresh token")
	}

	return nil
}
