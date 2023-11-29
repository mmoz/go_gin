package user

type (
	UserProfile struct {
		Username string `json:"username"`
		Role     string `json:"role"`
	}

	CreateUserReq struct {
		Username     string `json:"username"`
		Password     string `json:"password"`
		Role         string `json:"role"`
		RefreshToken string `json:"refresh_token,omitempty"`
	}
)
