package usecase

import (
	"github.com/pkg/errors"
	"go_clearn_architect/internal/domain/auth/presenter"
	"go_clearn_architect/internal/domain/auth/repository"
	"go_clearn_architect/pkg/utils"
)

func SignUp(signUpRequest *presenter.SignUpRequest, auth repository.IAuth) (presenter.SignUpResponse, error) {
	if _, err := auth.GetUserByEmail(signUpRequest.Email); err == nil {
		return presenter.SignUpResponse{}, errors.New("User is already register")
	}
	user, err := auth.CreateUser(signUpRequest)
	if err != nil {
		return presenter.SignUpResponse{}, err
	}
	accessToken, err := utils.GenerateJWT(user.ID, user.Email, user.Username)
	if err != nil {
		return presenter.SignUpResponse{}, errors.New("Can not generate access token")
	}
	return presenter.SignUpResponse{
		AccessToken: accessToken,
	}, nil
}
