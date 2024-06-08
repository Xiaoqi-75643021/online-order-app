package model

import "gorm.io/gorm"

type Cart struct {
	CartID uint `gorm:"column:id;primaryKey;autoIncrement"` // 购物车ID
	UserID uint `gorm:"column:user_id;index"`               // 用户ID
	gorm.Model
}

func (*Cart) TableName() string {
	return "cart"
}

type CartItem struct {
	CartItemID    uint    `gorm:"column:id;primaryKey;autoIncrement"`     // 购物详情ID
	CartID        uint    `gorm:"column:cart_id;index"`                   // 购物车ID，索引
	DishID        uint    `gorm:"column:dish_id;index"`                   // 菜品ID，索引
	Quantity      int     `gorm:"column:quantity;type:int"`               // 菜品数量
	Price         float64 `gorm:"column:price;type:decimal(10,2)"`        // 菜品的单价
	Specification string  `gorm:"column:specification;type:varchar(255)"` // 菜品规格
	// Note       string  `gorm:"column:note;type:text"`              // 备注(不加辣，不要酸黄瓜...)
	gorm.Model
}

func (*CartItem) TableName() string {
	return "cart_item"
}
