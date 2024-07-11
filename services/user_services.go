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

func (u *UserService) GetUserService(status bool) ([]*internal.User, error) {
	var users []*internal.User

	if err := u.db.Where("status = ?", status).Find(&users).Error; err != nil {
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

func (u *UserService) UpdateUserService(name string, status bool, id int) (*internal.User, error) {

	var user *internal.User

	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	user.Name = name
	user.Status = status

	if err := u.db.Save(&user).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return user, nil
}

func (u *UserService) DeleteUserService(id int64) error {

	var user *internal.User

	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}

	if err := u.db.Where("id = ?", id).Delete(&user).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
