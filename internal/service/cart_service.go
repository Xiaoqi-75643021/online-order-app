package service

import (
	"online-ordering-app/internal/database"
	"online-ordering-app/internal/model"
	"online-ordering-app/internal/repository"

	"gorm.io/gorm"
)

func AddToCart(userID, dishID uint) (int, error) {
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	cart, err := repository.FindCartByUserID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			cart = &model.Cart{UserID: userID}
			if err := repository.CreateCart(tx, cart); err != nil {
				tx.Rollback()
				return -1, err
			}
		} else {
			tx.Rollback()
			return -1, err
		}
	}

	dish, err := repository.FindDishByID(dishID)
	if err != nil {
		tx.Rollback()
		return int(cart.CartID), err
	}

	cartItem, err := repository.FindCartItemByCartIDAndDishID(cart.CartID, dishID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			cartItem = &model.CartItem{
				CartID:   cart.CartID,
				DishID:   dishID,
				Quantity: 1,
				Price:    dish.Price,
			}
			if err := repository.SaveCartItem(tx, cartItem); err != nil {
				tx.Rollback()
				return -1, err
			}
		} else {
			tx.Rollback()
			return int(cart.CartID), err
		}
	} else {
		cartItem.Quantity += 1
		if err := repository.SaveCartItem(tx, cartItem); err != nil {
			tx.Rollback()
			return -1, err
		}
	}

	return int(cart.CartID), tx.Commit().Error
}

func RemoveFromCart(userID, cartItemID uint) error {
	cart, err := repository.FindCartByUserID(userID)
	if err != nil {
		return err
	}

	if err := repository.DeleteCartItem(cartItemID); err != nil {
		return err
	}

	count, err := repository.CountCartItems(cart.CartID)
	if err != nil {
		return err
	}

	tx := database.DB.Begin()
	if count == 0 {
		if err := repository.DeleteCart(tx, cart.CartID); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func ClearCart(cartID uint) error {
	tx := database.DB.Begin()
	if err := repository.DeleteCart(tx, cartID); err !=nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func ListCartItemsByCartID(cartID uint) ([]*model.CartItem, error) {
	return repository.ListCartItemsByCartID(cartID)
}
