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

func DeleteCart(id uint) error {
	tx := database.DB.Begin()

	if err := tx.Where("cart_id = ?", id).Delete(&model.CartItem{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&model.Cart{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func SaveCartItem(cartItem *model.CartItem) error {
	return database.DB.Save(cartItem).Error
}

func FindCartItemByCartIDAndDishID(cartID, dishID uint) (*model.CartItem, error) {
	cartItem := new(model.CartItem)
	err := database.DB.Where("cart_id = ? and dish_id = ?", cartID, dishID).First(&cartItem).Error
	return cartItem, err
}

func DeleteCartItem(cartItemID uint) error {
	if err := database.DB.Delete(&model.CartItem{}, cartItemID).Error; err != nil {
		return err
	}
	return nil
}

func CountCartItems(cartID uint) (int64, error) {
	var count int64
	database.DB.Model(&model.CartItem{}).Where("cart_id = ?", cartID).Count(&count)
	return count, nil
}

func ListCartItemsByCartID(cartID uint) ([]*model.CartItem, error) {
	var cartItems []*model.CartItem
	err := database.DB.Where("cart_id = ?", cartID).Find(&cartItems).Error
	return cartItems, err
}
