package repository

import (
	"online-ordering-app/internal/database"
	"online-ordering-app/internal/model"
)

func CreateCart(cart *model.Cart) error {
	return database.DB.Create(cart).Error
}

func FindCartByID(cartID uint) (*model.Cart, error) {
	cart := new(model.Cart)
	err := database.DB.Where("id = ?", cartID).First(&cart).Error
	return cart, err
}

func FindCartByUserID(userID uint) (*model.Cart, error) {
	cart := new(model.Cart)
	err := database.DB.Where("user_id = ?", userID).First(&cart).Error
	return cart, err
}

func UpdateCart(cart *model.Cart) error {
	return database.DB.Save(cart).Error
}

func DeleteCart(id uint) error {
	return database.DB.Delete(&model.Cart{}, id).Error
}
