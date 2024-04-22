package service

import (
	"online-ordering-app/internal/repository"
	"online-ordering-app/internal/utils"
)

func UpdateAdminName(adminID uint, newAdminName string) error {
	admin, err := repository.FindAdminByID(adminID)
	if err != nil {
		return err
	}
	if _, err := repository.FindAdminByName(newAdminName); err == nil {
		return ErrUserAlreadyExists
	}
	admin.Name = newAdminName
	return repository.UpdateAdmin(admin)
}

func UpdateAdminPassword(adminID uint, newPassword string) error {
	admin, err := repository.FindAdminByID(adminID)
	if err != nil {
		return err
	}
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}
	admin.Password = hashedPassword
	return repository.UpdateAdmin(admin)
}
