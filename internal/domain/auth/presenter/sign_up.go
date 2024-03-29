package presenter

type (
	SignUpRequest struct {
		Username string `json:"username" validate:"required,gte=0,lte=130"`
		Password string `json:"password" validate:"required,gte=0,lte=16"`
		Email    string `json:"email" validate:"required,email"`
	}
	SignUpResponse struct {
		AccessToken string `json:"access_token"`
	}
)
