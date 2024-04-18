package model

import "gorm.io/gorm"

type Cart struct {
	CartID uint   `gorm:"column:id;primaryKey;autoIncrement"` // 购物车ID
	UserID uint   `gorm:"column:user_id;index"`               // 用户ID
	Note   string `gorm:"column:note;type:text"`              // 备注(不加辣，不要酸黄瓜...)
	gorm.Model
}

func (*Cart) TableName() string {
	return "cart"
}
