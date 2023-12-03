package user

type (
	UserProfile struct {
		Username string `json:"username"`
		Role     string `json:"role,omitempty"`
	}

	CreateUserReq struct {
		ID           string `json:"id"`
		Username     string `json:"username" validate:"required"`
		Password     string `json:"password" validate:"required"`
		Role         string `json:"role" validate:"required"`
		RefreshToken string `json:"refresh_token,omitempty"`
	}
)
