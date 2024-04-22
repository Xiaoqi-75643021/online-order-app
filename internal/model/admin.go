package model

import "gorm.io/gorm"

type Admin struct {
	AdminID  uint   `gorm:"column:id;primaryKey;autoIncrement"`        // 管理员ID，主键，自增
	Name     string `gorm:"column:name;type:varchar(100);uniqueIndex"` // 管理员名，唯一，索引
	Password string `gorm:"column:password;type:varchar(100)"`         // 管理员密码
	gorm.Model
}

func (*Admin) TableName() string {
	return "admin"
}
