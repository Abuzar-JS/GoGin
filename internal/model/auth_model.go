package model

type UserAuth struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Email    string `json:json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=8,max=200"`
}

func (UserAuth) TableName() string {
	return "userAuth"
}
