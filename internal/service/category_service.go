package service

import (
	"online-ordering-app/internal/model"
	"online-ordering-app/internal/repository"
)

func Addategory(categoryName string, parentID *uint) error {
	_, err := repository.FindCategoryByName(categoryName)
	if err == nil {
		return ErrCategoryAlreadyExists
	}
	category := &model.Category{
		Category: categoryName,
		ParentID: parentID,
	}
	return repository.CreateCategory(category)
}

func RemoveCategory(categoryID uint) error {
	return repository.DeleteCategory(categoryID)
}

func ListCategories(page, pageSize int) ([]*model.Category, error) {
	return repository.ListCategories(page, pageSize)
}

func QueryCategoryByID(categoryID uint) (*model.Category, error) {
	return repository.FindCategoryByID(categoryID)
}