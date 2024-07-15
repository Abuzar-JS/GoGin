package model

type UserAuth struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (UserAuth) TableName() string {
	return "userAuth"
}
