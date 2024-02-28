package controllers

import (
	"erdemkayacomtr/services"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	GetCategories(c *gin.Context)
	GetCategoryById(c *gin.Context)
	CreateCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
}

type CategoryControllerImpl struct {
	categoryService services.CategoryService
}

func (cat CategoryControllerImpl) GetCategories(c *gin.Context) {
	cat.categoryService.GetCategories(c)
}

func (cat CategoryControllerImpl) GetCategoryById(c *gin.Context) {
	cat.categoryService.GetCategoryById(c)
}

func (cat CategoryControllerImpl) CreateCategory(c *gin.Context) {
	cat.categoryService.CreateCategory(c)
}
func (cat CategoryControllerImpl) UpdateCategory(c *gin.Context) {
	cat.categoryService.UpdateCategory(c)
}
func (cat CategoryControllerImpl) DeleteCategory(c *gin.Context) {
	cat.categoryService.DeleteCategory(c)
}

func CategoryControllerInit(categoryService services.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		categoryService: categoryService,
	}
}
