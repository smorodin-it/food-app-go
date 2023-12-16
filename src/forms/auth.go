package forms

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
	return nil
}
