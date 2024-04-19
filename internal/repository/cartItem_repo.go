package repository

import (
	"online-ordering-app/internal/database"
	"online-ordering-app/internal/model"
)

func CreateCartItem(cartItem *model.CartItem) error {
	return database.DB.Create(cartItem).Error
}

func FindCartItemByID(cartItemID uint) (*model.CartItem, error) {
	cartItem := new(model.CartItem)
	err := database.DB.Where("id = ?", cartItemID).First(&cartItem).Error
	return cartItem, err
}

func FindCartItemsByCartID(cartID uint) ([]*model.CartItem, error) {
	var cartItems []*model.CartItem
	err := database.DB.Where("cart_id = ?", cartID).Find(&cartItems).Error
	return cartItems, err
}

func UpdateCartItem(cartItem *model.CartItem) error {
	return database.DB.Save(cartItem).Error
}

func DeleteCartItem(id uint) error {
	return database.DB.Delete(&model.CartItem{}, id).Error
}

func DeleteCartItemsByCartID(cartID uint) error {
	return database.DB.Where("cart_id = ?", cartID).Delete(&model.CartItem{}).Error
}
