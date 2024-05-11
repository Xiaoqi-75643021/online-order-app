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
	OrderItemID uint `gorm:"column:id;primaryKey;autoIncrement"`
	OrderID uint `gorm:"column:order_id;index"`
	gorm.Model
}