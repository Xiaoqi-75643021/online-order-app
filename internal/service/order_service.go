package service

import (
	"errors"
	"online-ordering-app/internal/database"
	"online-ordering-app/internal/model"
	"online-ordering-app/internal/repository"
)

func SubmitOrder(cartID, userID uint, orderType, note, payMethod string) (*model.Order, error) {
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

	totalamount := calculateTotalAmount(cartItems)
	if payMethod == "余额支付" {
		user, _ := repository.FindUserByID(userID)
		if totalamount > user.Balance {
			return nil, errors.New("余额不足")
		} else {
			user.Balance -= totalamount
			repository.UpdateUser(user)
		}
	} else if payMethod != "微信支付" {
		return nil, errors.New("不支持该支付方式")
	}


	order := &model.Order{
		UserID:      userID,
		TotalAmount: totalamount,
		Status:      "已支付",
		Type:        orderType,
		Note:        note,
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

func QueryOrderItemsByOrderID(orderID uint) ([]*model.OrderItem, error) {
	orderItems, err := repository.ListOrderItemsByOrderID(orderID)
	if err != nil {
		return nil, err
	}
	return orderItems, nil
}

func RemoveOrderByOrderID(orderID uint) error {
	if err := repository.DeleteOrder(orderID); err != nil {
		return err
	}
	return nil
}

func RefundByOrderID(userID, orderID uint) error {
	user, err := repository.FindUserByID(userID)
	if err != nil {
		return err
	}

	order, err := repository.FindOrderByID(orderID)
	if err != nil {
		return err
	}
	if err := repository.DeleteOrder(orderID); err != nil {
		return err
	}

	orderAmount := order.TotalAmount
	user.Balance += orderAmount

	err = repository.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func SubmitRefundRequestByOrderID(orderID uint) error {
	order, err := repository.FindOrderByID(orderID)
	if err != nil {
		return err
	}
	if order.Status == "退款中" {
		return errors.New("订单正在退款")
	}
	order.Status = "退款中"
	if err := repository.SaveOrder(order); err != nil {
		return err
	}

	return nil
}

func SubmitCommentByOrderID(orderID uint, comment string) error {
	order, err := repository.FindOrderByID(orderID)
	if err != nil {
		return err
	}
	order.Comment = comment
	if err := repository.SaveOrder(order); err != nil {
		return err
	}

	return nil
}