package model

import (
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	CouponID uint      `gorm:"column:id;primaryKey;autoIncrement"` // 优惠券ID，主键，自增
	Name     string    `gorm:"column:name;type:varchar(50)"`       // 优惠劵名称
	Amount   float64   `gorm:"column:amount;type:decimal(10,2)"`   // 用券价格
	DishID   uint      `gorm:"column:dish_id;index"`               // 商品ID，索引
	Quantity int       `gorm:"column:quantity;type:int"`           // 商品数量
	StartDay time.Time `gorm:"column:start_day;type:date"`         // 开始有效期
	EndDay   time.Time `gorm:"column:end_day;type:date"`           // 结束有效期
	gorm.Model
}

func (c *Coupon) TableName() string {
	return "coupon"
}

type UserCoupon struct {
	UserCouponID uint   `gorm:"column:id;primaryKey;autoIncrement"` // 用户优惠券ID，主键，自增
	UserID       uint   `gorm:"column:user_id;index"`               // 用户ID，索引
	CartID       uint   `gorm:"column:cart_id;index"`               // 购物车ID，索引
	CouponID     uint   `gorm:"column:coupon_id;index"`             // 优惠券ID，索引
	Status       string `gorm:"column:status;type:varchar(50)"`     // 状态，例如：未使用、已使用
	gorm.Model
}

func (uc *UserCoupon) TableName() string {
	return "user_coupon"
}
