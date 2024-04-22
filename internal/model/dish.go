package model

import "gorm.io/gorm"

type Dish struct {
	DishId      uint    `gorm:"column:id;primaryKey;autoIncrement"`        // 菜品ID，主键，自增
	Name        string  `gorm:"column:name;type:varchar(100);uniqueIndex"` // 菜品名称
	Description string  `gorm:"column:description;type:text"`              // 菜品描述
	Price       float64 `gorm:"column:price;type:decimal(10,2)"`           // 菜品价格
	CategoryID  uint    `gorm:"column:category_id;index"`                  // 分类ID，索引
	IsPopular   bool    `gorm:"column:is_popular;default:false"`           // 是否为热门菜品，默认为false
	gorm.Model
}

func (d *Dish) TableName() string {
	return "dish"
}
