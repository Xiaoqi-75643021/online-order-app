package model

import "gorm.io/gorm"

type Order struct {
	OrderID     uint    `gorm:"column:id;primaryKey;autoIncrement"`     // 订单ID
	UserID      uint    `gorm:"column:user_id;index"`                   // 用户ID
	TotalAmount float64 `gorm:"column:total_amount;type:decimal(10,2)"` // 总金额
	Status      string  `gorm:"column:status;type:varchar(50)"`         // 未支付，制作中，已备好，完成
	Type        string  `gorm:"column:type;type:varchar(50)"`           // 堂食or外带
	gorm.Model
}

func (*Order) TableName() string {
	return "order"
}

type OrderItem struct {
	OrderItemID uint   `gorm:"column:id;primaryKey;autoIncrement"` // 订单项ID
	OrderID     uint   `gorm:"column:order_id;index"`              // 订单ID
	DishID      uint   `gorm:"column:dish_id;index"`               // 菜品ID，索引
	Quantity    int    `gorm:"column:quantity;type:int"`           // 菜品数量
	Note        string `gorm:"column:note;type:text"`              // 备注(不加辣，不要酸黄瓜...)
}

func (*OrderItem) TableName() string {
	return "order_item"
}
