package model

import "gorm.io/gorm"

type Category struct {
	CategoryID uint   `gorm:"column:id;primaryKey;autoIncrement"` // 分类ID，主键，自增
	Category   string `gorm:"column:category;type:varchar(100)"`  // 分类名(甜品/小食/炸鸡/汉堡)
	ParentID   *uint  `gorm:"column:parent_id;index"`             // 指向父级ID，如果为null表示顶级分类
	gorm.Model
}

func (*Category) TableName() string {
	return "category"
}
