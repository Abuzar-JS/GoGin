package services

import (
	internal "GinMod/internal/model"
	"fmt"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func (u *UserService) InitService(db *gorm.DB) {
	u.db = db
}

func (u *UserService) GetUserService() string {
	return "Get User Request "
}

func (u *UserService) CreateUserService(name string, status bool) (*internal.User, error) {
	if u.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	user := &internal.User{
		Name:   name,
		Status: status,
	}
	if err := u.db.Create(user).Error; err != nil {
		fmt.Print(err)
		return nil, err
	}
	return user, nil
}
