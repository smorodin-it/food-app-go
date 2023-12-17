package forms

import "errors"

type FormAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (f FormAuth) Validate() error {
	return nil
}

type RefreshForm struct {
	RefreshToken string `json:"refreshToken"`
}

func (f RefreshForm) Validate() error {
	if f.RefreshToken == "" {
		return errors.New("bad refresh token")
	}

	return nil
}

type AccessTokenForm struct {
	Token string `json:"token"`
}
