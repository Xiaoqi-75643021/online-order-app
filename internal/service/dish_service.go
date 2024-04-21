package service

import (
	"online-ordering-app/internal/model"
	"online-ordering-app/internal/repository"
)

func SearchDishes(keyword string, page, pageSize int) ([]*model.Dish, error) {
	return repository.ListDishesByKeyword(keyword, page, pageSize)
}

func GetDishesByCategory(categoryId uint, page, pageSize int) ([]*model.Dish, error) {
	return repository.ListDishesByCategory(categoryId, page, pageSize)
}

func GetPopularDishes() ([]*model.Dish, error) {
	return repository.ListPopularDishes()
}