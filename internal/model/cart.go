package model

import "gorm.io/gorm"

type Cart struct {
	CartID uint   `gorm:"column:id;primaryKey;autoIncrement"` // 购物车ID
	UserID uint   `gorm:"column:user_id;index"`               // 用户ID
	gorm.Model
}

func (*Cart) TableName() string {
	return "cart"
}
