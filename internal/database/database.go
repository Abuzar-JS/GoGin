package internal

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5400
	user     = "postgres"
	password = "1234"
	dbName   = "DbGin"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host = %s port =%d user =%s password=%s dbname=%s sslmode= disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	if err != nil {
		fmt.Println("Unable to connect to Database")
	}

	return db
}
