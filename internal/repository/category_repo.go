package repository

import (
	"online-ordering-app/internal/database"
	"online-ordering-app/internal/model"
)

func CreateCategory(category *model.Category) error {
	return database.DB.Create(category).Error
}

func FindCategoryByID(categoryID uint) (*model.Category, error) {
	category := new(model.Category)
	err := database.DB.Where("id = ?", categoryID).First(&category).Error
	return category, err
}

func FindCategoryByName(name string) (*model.Category, error) {
	category := new(model.Category)
	err := database.DB.Where("category = ?", name).First(&category).Error
	return category, err
}

func UpdateCategory(category *model.Category) error {
	return database.DB.Save(category).Error
}

func DeleteCategory(id uint) error {
	return database.DB.Delete(&model.Category{}, id).Error
}

func ListCategories(page, pageSize int) ([]*model.Category, error) {
	var categories []*model.Category
	err := database.DB.Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&categories).Error
	return categories, err
}

func ListSubCategories(parentID uint, page, pageSize int) ([]*model.Category, error) {
	var subCategories []*model.Category
	err := database.DB.Where("parent_id = ?", parentID).Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&subCategories).Error
	return subCategories, err
}
