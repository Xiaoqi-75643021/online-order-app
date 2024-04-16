package service

import (
	"online-ordering-app/internal/repository"
	"online-ordering-app/internal/utils"
)

func UpdateUsername(userID uint, newUsername string) error {
	user, err := repository.FindUserByID(userID)
	if err != nil {
		return err
	}
	if _, err := repository.FindUserByName(newUsername); err == nil {
		return ErrUserAlreadyExists
	}
	user.Username = newUsername
	return repository.UpdateUser(user)
}

func Updatepassword(userID uint, newPassword string) error {
	user, err := repository.FindUserByID(userID)
	if err != nil {
		return err
	}
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return repository.UpdateUser(user)
}
