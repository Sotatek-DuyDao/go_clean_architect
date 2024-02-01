package repository

import (
	"github.com/pkg/errors"
	"go_clearn_architect/internal/adapter/mysql"
	"go_clearn_architect/internal/adapter/mysql/model"
	"go_clearn_architect/internal/domain/auth/presenter"
	"go_clearn_architect/pkg/utils"
	"gorm.io/gorm"
)

type Auth struct{}

func (auth Auth) GetUserById(userId uint) (model.User, error) {
	db := mysql.Connect()
	var user = model.User{}
	db.First(&user, userId)
	return user, nil
}

func (auth Auth) GetUserByEmail(userEmail string) (model.User, error) {
	db := mysql.Connect()
	var user = model.User{
		Email: userEmail,
	}
	if err := db.Where(&user).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, err
	}
	return user, nil
}

func (auth Auth) CreateUser(userSignUp *presenter.SignUpRequest) (model.User, error) {
	db := mysql.Connect()
	hashPassword, _ := utils.Hash(userSignUp.Password)
	var user = model.User{
		Email:    userSignUp.Email,
		Password: hashPassword,
		Username: userSignUp.Username,
	}
	if err := db.Create(&user); err.Error != nil {
		return model.User{}, err.Error
	}
	return user, nil
}
