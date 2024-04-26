package model

import "gorm.io/gorm"

type Coupon struct {
	CouponID       uint    `gorm:"column:id;primaryKey;autoIncrement"`        // 优惠券ID，主键，自增
	Description    string  `gorm:"column:description;type:text"`              // 优惠劵描述
	DiscountAmount float64 `gorm:"column:discount_amount;type:decimal(10,2)"` // 折扣金额
	Validity       string  `gorm:"column:validity;type:varchar(50)"`          // 有效期
	gorm.Model
}

func (c *Coupon) TableName() string {
	return "coupon"
}

type UserCoupon struct {
	UserCouponID uint   `gorm:"column:id;primaryKey;autoIncrement"` // 用户优惠券ID，主键，自增
	UserID       uint   `gorm:"column:user_id;index"`               // 用户ID，索引
	CouponID     uint   `gorm:"column:coupon_id;index"`             // 优惠券ID，索引
	Status       string `gorm:"column:status;type:varchar(50)"`     // 状态，例如：未使用、已使用
	gorm.Model
}

func (uc *UserCoupon) TableName() string {
	return "user_coupon"
}
