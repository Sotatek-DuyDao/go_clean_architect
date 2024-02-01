package repository

import (
	"go_clearn_architect/internal/adapter/mysql/model"
	"go_clearn_architect/internal/domain/auth/presenter"
)

type IAuth interface {
	GetUserById(userId uint) (model.User, error)
	GetUserByEmail(userEmail string) (model.User, error)
	CreateUser(userSignUp *presenter.SignUpRequest) (model.User, error)
}
