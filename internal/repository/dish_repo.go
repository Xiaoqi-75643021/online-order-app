package repository

import (
	"online-ordering-app/internal/database"
	"online-ordering-app/internal/model"
)

func CreateDish(dish *model.Dish) error {
	return database.DB.Create(dish).Error
}

func FindDishByID(dishID uint) (*model.Dish, error) {
	dish := new(model.Dish)
	err := database.DB.Where("id = ?", dishID).First(&dish).Error
	return dish, err
}

func FindDishByName(name string) (*model.Dish, error) {
	dish := new(model.Dish)
	err := database.DB.Where("name = ?", name).First(&dish).Error
	return dish, err
}

func UpdateDish(dish *model.Dish) error {
	return database.DB.Save(dish).Error
}

func DeleteDish(id uint) error {
	return database.DB.Delete(&model.Dish{}, id).Error
}

func ListDishes(page, pageSize int) ([]*model.Dish, error) {
	var dishes []*model.Dish
	err := database.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dishes).Error
	return dishes, err
}

func ListDishesByCategory(categoryID uint, page, pageSize int) ([]*model.Dish, error) {
	var dishes []*model.Dish
	err := database.DB.Where("category_id = ?", categoryID).Offset((page - 1) * pageSize).Limit(pageSize).Find(&dishes).Error
	return dishes, err
}
