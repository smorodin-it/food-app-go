package forms

type FormAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (f FormAuth) Validate() error {
	return nil
}
