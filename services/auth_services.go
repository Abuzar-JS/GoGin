package services

import (
	internal "GinMod/internal/model"
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
		return nil, errors.New("Email can't be null")
	}

	if password == nil {
		return nil, errors.New("Password can't be empty")
	}

	var user internal.UserAuth

	if err := a.db.Where("email=?", email).Where("password=?").Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *AuthService) Register(email *string, password *string) (*internal.UserAuth, error) {
	if email == nil {
		return nil, errors.New("email can't be null")
	}

	if password == nil {
		return nil, errors.New("password can't be empty")
	}

	var user internal.UserAuth

	user.Email = *email
	user.Password = *password

	if err := a.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
