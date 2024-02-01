package usecase

import (
	"github.com/pkg/errors"
	"go_clearn_architect/internal/domain/auth/presenter"
	"go_clearn_architect/internal/domain/auth/repository"
)

func SignUp(signUpRequest *presenter.SignUpRequest, auth repository.IAuth) (presenter.SignUpResponse, error) {
	if _, err := auth.GetUserByEmail(signUpRequest.Email); err == nil {
		return presenter.SignUpResponse{}, errors.New("User is already register")
	}
	user, err := auth.CreateUser(signUpRequest)
	if err != nil {
		return presenter.SignUpResponse{}, err
	}
	return presenter.SignUpResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
