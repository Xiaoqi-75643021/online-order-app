package repository

import (
	"online-ordering-app/internal/database"
	"online-ordering-app/internal/model"
)

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
