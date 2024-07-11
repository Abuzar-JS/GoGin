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

type User struct {
	Id   int
	Name string
}

func (u *UserService) GetUserService() ([]*internal.User, error) {
	var users []*internal.User

	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil

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
