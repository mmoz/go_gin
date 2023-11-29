package auth

type (
	Credential struct {
		Username     string `json:"username"`
		Role         string `json:"role"`
		RefreshToken string `json:"refreshtoken"`
	}
)
