package auth

type (
	CredentialReq struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	RefreshTokenReq struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}

	CredentialRes struct {
		Username     string `json:"username"`
		Password     string `json:"password,omitempty"`
		Role         string `json:"role"`
		RefreshToken string `json:"RefreshToken"`
		AccessToken  string `json:"AccessToken"`
	}
)
