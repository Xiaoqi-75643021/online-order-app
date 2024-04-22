package repository

import (
	"online-ordering-app/internal/database"
	"online-ordering-app/internal/model"
	"online-ordering-app/internal/utils"

	"gorm.io/gorm"
)

func CreateUser(user *model.User) error {
	return database.DB.Create(user).Error
}

func FindUserByName(username string) (*model.User, error) {
	user := new(model.User)
	err := database.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

func FindUserByID(userID uint) (*model.User, error) {
	user := new(model.User)
	err := database.DB.Where("id = ?", userID).First(&user).Error
	return user, err
}

func UpdateUser(user *model.User) error {
	return database.DB.Save(user).Error
}

func DeleteUser(id uint) error {
	return database.DB.Delete(&model.User{}, id).Error
}

func AuthenticateUser(username, password string) (bool, error) {
	user, err := FindUserByName(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return false, nil
	}
	return true, nil
}