package internal

type User struct {
	Id     int    `gorm:"primaryKey"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}
