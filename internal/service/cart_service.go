package service

import (
	"online-ordering-app/internal/model"
	"online-ordering-app/internal/repository"

	"gorm.io/gorm"
)

func AddToCart(userID, dishID uint) (int, error) {
	cart, err := repository.FindCartByUserID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			cart = &model.Cart{UserID: userID}
			if err := repository.CreateCart(cart); err != nil {
				return -1, err
			}
		} else {
			return -1, err
		}
	}

	dish, err := repository.FindDishByID(dishID)
	if err != nil {
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
		} else {
			return int(cart.CartID), err
		}
	} else {
		cartItem.Quantity += 1
	}
	return int(cart.CartID), repository.SaveCartItem(cartItem)
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

	if count == 0 {
		return repository.DeleteCart(cart.CartID)
	}

	return nil
}

func ClearCart(cartID uint) error {
	return repository.DeleteCart(cartID)
}

func ListCartItemsByCartID(cartID uint) ([]*model.CartItem, error) {
	return repository.ListCartItemsByCartID(cartID)
}
