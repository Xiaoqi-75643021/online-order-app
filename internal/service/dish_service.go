package service

import (
	"online-ordering-app/internal/model"
	"online-ordering-app/internal/repository"

	"github.com/mitchellh/mapstructure"
)

func AddDish(name string, price float64, category string) error {
	_, err := repository.FindDishByName(name)
	if err == nil {
		return ErrDishAlreadyExists
	}
	categoryStruct, err := repository.FindCategoryByName(category)
	categoryID := categoryStruct.CategoryID
	if err != nil {
		return err
	}
	dish := &model.Dish{
		Name:       name,
		Price:      price,
		CategoryID: categoryID,
	}
	return repository.CreateDish(dish)
}

func DeleteDish(id uint) error {
	return repository.DeleteDish(id)
}

func UpdateDish(dishID uint, dishUpdate map[string]any) error {
	dish, err := repository.FindDishByID(dishID)
	if err != nil {
		return err
	}
	decoderConfig := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result: &dish,
		TagName: "json",
	}

	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err != nil {
		return err
	}

	err = decoder.Decode(dishUpdate)
	if err != nil {
		return err
	}

	err = repository.UpdateDish(dish)
	if err != nil {
		return err
	}

	return nil
}

func SearchDishes(keyword string, page, pageSize int) ([]*model.Dish, error) {
	return repository.ListDishesByKeyword(keyword, page, pageSize)
}

func GetDishesByCategory(categoryId uint, page, pageSize int) ([]*model.Dish, error) {
	return repository.ListDishesByCategory(categoryId, page, pageSize)
}

func GetPopularDishes() ([]*model.Dish, error) {
	return repository.ListPopularDishes()
}

func ListDishes(page, pageSize int) ([]*model.Dish, error) {
	return repository.ListDishes(page, pageSize)
}