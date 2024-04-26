package repository

import (
	"online-ordering-app/internal/database"
	"online-ordering-app/internal/model"
)

func CreateCoupon(coupon *model.Coupon) error {
	return database.DB.Create(coupon).Error
}

func FindCouponByID(couponID uint) (*model.Coupon, error) {
	coupon := new(model.Coupon)
	err := database.DB.Where("id = ?", couponID).First(&coupon).Error
	return coupon, err
}

func UpdateCoupon(coupon *model.Coupon) error {
	return database.DB.Save(coupon).Error
}

func DeleteCoupon(id uint) error {
	return database.DB.Delete(&model.Coupon{}, id).Error
}

func ListCoupons(page, pageSize int) ([]*model.Coupon, error) {
	var coupons []*model.Coupon
	err := database.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&coupons).Error
	return coupons, err
}

func CreateUserCoupon(userCoupon *model.UserCoupon) error {
	return database.DB.Create(userCoupon).Error
}

func FindUserCouponByID(userCouponID uint) (*model.UserCoupon, error) {
	userCoupon := new(model.UserCoupon)
	err := database.DB.Where("id = ?", userCouponID).First(&userCoupon).Error
	return userCoupon, err
}

func FindUserCouponsByUserID(userID uint) ([]*model.UserCoupon, error) {
	var userCoupons []*model.UserCoupon
	err := database.DB.Where("user_id = ?", userID).Find(&userCoupons).Error
	return userCoupons, err
}

func UpdateUserCoupon(userCoupon *model.UserCoupon) error {
	return database.DB.Save(userCoupon).Error
}

func DeleteUserCoupon(id uint) error {
	return database.DB.Delete(&model.UserCoupon{}, id).Error
}
