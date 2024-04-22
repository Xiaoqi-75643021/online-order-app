package repository

import (
	"online-ordering-app/internal/database"
	"online-ordering-app/internal/model"
	"online-ordering-app/internal/utils"

	"gorm.io/gorm"
)

func CreateAdmin(admin *model.Admin) error {
	return database.DB.Create(admin).Error
}

func FindAdminByName(adminname string) (*model.Admin, error) {
	admin := new(model.Admin)
	err := database.DB.Where("name = ?", adminname).First(&admin).Error
	return admin, err
}

func FindAdminByID(adminID uint) (*model.Admin, error) {
	admin := new(model.Admin)
	err := database.DB.Where("id = ?", adminID).First(&admin).Error
	return admin, err
}

func UpdateAdmin(admin *model.Admin) error {
	return database.DB.Save(admin).Error
}

func DeleteAdmin(id uint) error {
	return database.DB.Delete(&model.Admin{}, id).Error
}

func AuthenticateAdmin(adminname, password string) (bool, error) {
	admin, err := FindAdminByName(adminname)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	if !utils.CheckPasswordHash(password, admin.Password) {
		return false, nil
	}
	return true, nil
}