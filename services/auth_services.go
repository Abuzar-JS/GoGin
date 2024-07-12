package services

import (
	"GinMod/internal/model"
	internal "GinMod/internal/model"
	"GinMod/utils"
	"errors"

	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func InitAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (a *AuthService) Login(email *string, password *string) (*internal.UserAuth, error) {
	if email == nil {
		return nil, errors.New("email can't be null")
	}

	if password == nil {
		return nil, errors.New("password can't be empty")
	}

	var user internal.UserAuth

	if err := a.db.Where("email=?", email).Find(&user).Error; err != nil {
		return nil, err
	}

	if utils.CheckPasswordHash(*password, user.Password) == false {
		return nil, errors.New("password not")
	}

	if user.Email == "" {
		return nil, errors.New("no user found with this email")
	}

	return &user, nil
}

func (a *AuthService) Register(email *string, password *string) (*model.UserAuth, error) {
	if email == nil {
		return nil, errors.New("email can't be null")
	}

	if password == nil {
		return nil, errors.New("password can't be empty")
	}

	hashedPwd, err := utils.HashPassword(*password)

	if err != nil {
		return nil, err
	}

	var user model.UserAuth

	user.Email = *email
	user.Password = hashedPwd

	if err := a.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
