package internal

type User struct {
	Id     int    `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}
