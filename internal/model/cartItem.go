package model

import "gorm.io/gorm"

type CartItem struct {
	CartItemID uint    `gorm:"column:id;primaryKey;autoIncrement"` // 购物详情ID
	CartID     uint    `gorm:"column:cart_id;index"`               // 购物者ID，索引
	DishID     uint    `gorm:"column:dish_id;index"`               // 菜品ID，索引
	Quantity   int     `gorm:"column:quantity;type:int"`           // 菜品数量
	Price      float64 `gorm:"column:price;type:decimal(10,2)"`    // 菜品的单价
	gorm.Model
}

func (*CartItem) TableName() string {
	return "cart_item"
}
