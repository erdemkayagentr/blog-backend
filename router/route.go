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
	}

	return router
}
