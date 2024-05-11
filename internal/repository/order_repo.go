package repository

import (
	"online-ordering-app/internal/database"
	"online-ordering-app/internal/model"

	"gorm.io/gorm"
)

func CreateOrder(tx *gorm.DB, order *model.Order) error {
	return tx.Create(order).Error
}

func FindOrderByID(orderID uint) (*model.Order, error) {
	order := new(model.Order)
	err := database.DB.Where("id = ?", orderID).First(&order).Error
	return order, err
}

func FindOrderByUserID(userID uint) (*model.Order, error) {
	order := new(model.Order)
	err := database.DB.Where("user_id = ?", userID).First(&order).Error
	return order, err
}

func DeleteOrder(id uint) error {
	tx := database.DB.Begin()

	if err := tx.Where("order_id = ?", id).Delete(&model.OrderItem{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&model.Order{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func SaveOrderItem(orderItem *model.OrderItem) error {
	return database.DB.Save(orderItem).Error
}

func FindOrderItemByOrderIDAndDishID(orderID, dishID uint) (*model.OrderItem, error) {
	orderItem := new(model.OrderItem)
	err := database.DB.Where("order_id = ? and dish_id = ?", orderID, dishID).First(&orderItem).Error
	return orderItem, err
}

func DeleteOrderItem(orderItemID uint) error {
	if err := database.DB.Delete(&model.OrderItem{}, orderItemID).Error; err != nil {
		return err
	}
	return nil
}

func CountOrderItems(orderID uint) (int64, error) {
	var count int64
	database.DB.Model(&model.OrderItem{}).Where("order_id = ?", orderID).Count(&count)
	return count, nil
}

func ListOrderItemsByOrderID(orderID uint) ([]*model.OrderItem, error) {
	var orderItems []*model.OrderItem
	err := database.DB.Where("order_id = ?", orderID).Find(&orderItems).Error
	return orderItems, err
}

func ConvertCartItemsToOrderItems(tx *gorm.DB, cartItems []*model.CartItem, orderID uint) error {
	for _, cartItem := range cartItems {
		orderItem := &model.OrderItem{
			OrderID:  orderID,
			DishID:   cartItem.DishID,
			Quantity: cartItem.Quantity,
			Note:     cartItem.Note,
		}
		if err := tx.Create(orderItem).Error; err != nil {
			return err
		}
	}
	return nil
}