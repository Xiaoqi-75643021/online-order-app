package service

import (
	"online-ordering-app/internal/database"
	"online-ordering-app/internal/model"
	"online-ordering-app/internal/repository"
)

func SubmitOrder(cartID, userID uint, orderType string) (*model.Order, error) {
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	cartItems, err := repository.ListCartItemsByCartID(cartID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	order := &model.Order{
		UserID:      userID,
		TotalAmount: calculateTotalAmount(cartItems),
		Status:      "已支付",
		Type:        orderType,
	}
	if err := repository.CreateOrder(tx, order); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := repository.ConvertCartItemsToOrderItems(tx, cartItems, order.OrderID); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := repository.DeleteCart(tx, cartID); err != nil {
		tx.Rollback()
		return nil, err
	}

	return order, tx.Commit().Error
}

func QueryOrdersByUserID(userID uint) ([]model.Order, error) {
	orders, err := repository.GetOrdersByUserID(userID)
	return orders, err
}

func calculateTotalAmount(cartItems []*model.CartItem) float64 {
	var total float64
	for _, item := range cartItems {
		total += float64(item.Quantity) * item.Price
	}
	return total
}