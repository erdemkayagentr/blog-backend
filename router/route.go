package router

import (
	"erdemkayacomtr/config"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		categories := api.Group("/categories")
		categories.GET("", init.CategoryController.GetCategories)
		categories.POST("", init.CategoryController.CreateCategory)
		categories.GET("/:id", init.CategoryController.GetCategoryById)
		categories.PUT("/:id", init.CategoryController.UpdateCategory)
		categories.DELETE("/:id", init.CategoryController.DeleteCategory)
	}

	return router
}
