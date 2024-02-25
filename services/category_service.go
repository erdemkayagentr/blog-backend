package services

import (
	"erdemkayacomtr/constant"
	"erdemkayacomtr/pkg"
	"erdemkayacomtr/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CategoryService interface {
	GetCategories(c *gin.Context)
	// GetCategoryById() (c *gin.Context)
	// CreateCategory() (c *gin.Context)
	// UpdateCategory() (c *gin.Context)
	// DeleteCategory() (c *gin.Context)
}

type CategoryServiceImpl struct {
	categoryRepository repositories.CategoryRepository
}

func (cat CategoryServiceImpl) GetCategories(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get all data categories")

	data, err := cat.categoryRepository.GetCategories()
	if err != nil {
		log.Error("Happened Error when find all category data. Error: ", err)
		pkg.PanicException(constant.UnkonwnError)
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data, ""))
}

func CategoryServiceInit(categoryRepository repositories.CategoryRepository) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		categoryRepository: categoryRepository,
	}
}
