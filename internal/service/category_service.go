package service

import (
	"online-ordering-app/internal/model"
	"online-ordering-app/internal/repository"
)

func CreateCategory(categoryName string, parentID *uint) error {
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

func DeleteCategory(categoryID uint) error {
	return repository.DeleteCategory(categoryID)
}

func ListCategories(page, pageSize int) ([]*model.Category, error) {
	return repository.ListCategories(page, pageSize)
}
