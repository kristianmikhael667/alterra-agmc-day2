package database

import (
	"main/config"
	"main/models"
)

// get all user
func GetUsers() (interface{}, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// get single user
func GetSingleUser(id string) (interface{}, error) {
	var users []models.User
	
	if err := config.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// delete user
func DeleteUser(id string) (interface{}, error) {

	var users []models.User

	if err := config.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Delete(&users, id).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func EditUser(id string, name string, email string, password string) (interface{}, error) {

	// var users []User
	var user []models.User

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Model(&user).Where("id = ?", id).Updates(models.User{Name: name, Email: email, Password: password}).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// create new user
func CreateUser(name string, email string, password string) (interface{}, error) {

	var users = []models.User{{Name: name, Email: email, Password: password}}

	if err := config.DB.Create(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
