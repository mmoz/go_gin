package user

type (
	UserProfile struct {
		Username string `json:"username"`
		Role     string `json:"role"`
	}

	CreateUserReq struct {
		Username     string `json:"username" validate:"required"`
		Password     string `json:"password" validate:"required"`
		Role         string `json:"role" validate:"required"`
		RefreshToken string `json:"refresh_token,omitempty"`
	}
)
