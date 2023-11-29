package auth

type (
	CredentialReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	RefreshTokenReq struct {
		RefreshToken string `json:"refresh_token"`
	}

	CredentialRes struct {
		Username     string `json:"username"`
		Role         string `json:"role"`
		RefreshToken string `json:"RefreshToken"`
		AccessToken  string `json:"AccessToken"`
	}
)
