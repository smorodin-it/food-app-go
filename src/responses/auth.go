package responses

type ResponseTokens struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
