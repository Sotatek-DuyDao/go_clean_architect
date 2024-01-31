package presenter

type (
	SignUpRequest struct {
		Username string `json:"username" validate:"required,gte=0,lte=130"`
		Password string `json:"password" validate:"required,gte=0,lte=16"`
		Email    string `json:"email" validate:"required,email"`
	}
)

type SignUpResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"password"`
}
