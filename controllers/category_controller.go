package controllers

import (
	"erdemkayacomtr/services"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	GetCategories(c *gin.Context)
}

type CategoryControllerImpl struct {
	categoryService services.CategoryService
}

func (cat CategoryControllerImpl) GetCategories(c *gin.Context) {
	cat.categoryService.GetCategories(c)
}

func CategoryControllerInit(categoryService services.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		categoryService: categoryService,
	}
}
