package service

import (
	"online-ordering-app/internal/database"
	"online-ordering-app/internal/model"
	"online-ordering-app/internal/repository"
)

func SubmitOrder(cartID, userID uint, orderType string) error {
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	cartItems, err := repository.ListCartItemsByCartID(cartID)
	if err != nil {
		tx.Rollback()
		return err
	}

	order := &model.Order{
		UserID:      userID,
		TotalAmount: calculateTotalAmount(cartItems),
		Status:      "未支付",
		Type:        orderType,
	}
	if err := repository.CreateOrder(tx, order); err != nil {
		tx.Rollback()
		return err
	}

	if err := repository.ConvertCartItemsToOrderItems(tx, cartItems, order.OrderID); err != nil {
		tx.Rollback()
		return err
	}

	if err := repository.DeleteCart(tx, cartID); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func calculateTotalAmount(cartItems []*model.CartItem) float64 {
	var total float64
	for _, item := range cartItems {
		total += float64(item.Quantity) * item.Price
	}
	return total
}