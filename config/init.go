package config

import (
	"erdemkayacomtr/controllers"
	"erdemkayacomtr/repositories"
	"erdemkayacomtr/services"
)

type Initialization struct {
	categoryRepository repositories.CategoryRepository
	categoryService    services.CategoryService
	CategoryController controllers.CategoryController
}

func NewInitializer(categoryRepository repositories.CategoryRepository,
	categoryService services.CategoryService,
	categoryController controllers.CategoryController) *Initialization {
	return &Initialization{
		categoryRepository: categoryRepository,
		categoryService:    categoryService,
		CategoryController: categoryController,
	}
}
