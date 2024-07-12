package internal

type UserAuth struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (UserAuth) TableName() string {
	return "userAuth"
}
