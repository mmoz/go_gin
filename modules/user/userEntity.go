package user

type (
	UserProfileEnt struct {
		ID            string `json:"id"`
		Username      string `json:"username"`
		Password      string `json:"password"`
		Role          string `json:"role" `
		RefreshToken  string `json:"refreshtoken"`
		IsTokenActive int    `json:"istokenactive"`
	}
)
