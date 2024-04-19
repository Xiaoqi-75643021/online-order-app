package model

import "gorm.io/gorm"

type User struct {
	UserID   uint    `gorm:"column:id;primaryKey;autoIncrement"`            // 用户ID，主键，自增
	Username string  `gorm:"column:username;type:varchar(100);uniqueIndex"` // 用户名，唯一，索引
	Password string  `gorm:"column:password;type:varchar(100)"`             // 用户密码（存储的是加密后的密码）
	Role     string  `gorm:"column:role;type:varchar(50)"`                  // 账户角色，区分普通用户和管理员
	Balance  float64 `gorm:"column:balance;type:decimal(10, 2)"`            // 余额
	gorm.Model
}

func (u *User) TableName() string {
	return "user"
}
