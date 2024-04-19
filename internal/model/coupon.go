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
