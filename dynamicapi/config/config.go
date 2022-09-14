package config

import (
	"fmt"
	"main/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// database connection
func InitDB() {
	// deklarasikan struct config & variable connectionString
	config := map[string]string{
		"DB_USERNAME": "root",
		"DB_PASSWORD": "",
		"DB_PORT":     "3306",
		"DB_HOST":     "127.0.0.1",
		"DB_NAME":     "alterramvc",
	}
	var err error
	connectionString := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config["DB_USERNAME"], config["DB_HOST"], config["DB_PORT"], config["DB_NAME"])

	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
}
