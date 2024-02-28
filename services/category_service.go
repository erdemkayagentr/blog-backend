package services

import (
	"erdemkayacomtr/constant"
	"erdemkayacomtr/domain/entities"
	"erdemkayacomtr/pkg"
	"erdemkayacomtr/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type CategoryService interface {
	GetCategories(c *gin.Context)
	GetCategoryById(c *gin.Context)
	CreateCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
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

func (cat CategoryServiceImpl) GetCategoryById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info(
		"start to execute get category by id",
	)

	data, err := cat.categoryRepository.GetCategoryById(uuid.MustParse(c.Param("id")))

	if err != nil {
		log.Error("Happened Error when find category by id. Error: ", err)
		pkg.PanicException(constant.UnkonwnError)
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data, ""))
}

func (cat CategoryServiceImpl) CreateCategory(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info(
		"start to execute create category",
	)

	var category entities.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		log.Error("Happened Error when bind json. Error: ", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := cat.categoryRepository.CreateCategory(&category)
	if err != nil {
		log.Error("Happened Error when create category. Error: ", err)
		pkg.PanicException(constant.UnkonwnError)
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data, ""))
}

func (cat CategoryServiceImpl) UpdateCategory(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info(
		"start to execute update category",
	)

	var category entities.Category
	catId := uuid.MustParse(c.Param("id"))

	if err := c.ShouldBindJSON(&category); err != nil {
		log.Error("Happened Error when bind json. Error: ", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := cat.categoryRepository.GetCategoryById(catId)
	if err != nil {
		log.Error("Happened Error when find category by id. Error: ", err)
		pkg.PanicException(constant.UnkonwnError)
	}

	data.Name = category.Name
	updatedData, err := cat.categoryRepository.UpdateCategory(&data)
	if err != nil {
		log.Error("Happened Error when update category. Error: ", err)
		pkg.PanicException(constant.UnkonwnError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, updatedData, ""))
}

func (cat CategoryServiceImpl) DeleteCategory(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info(
		"start to execute delete category",
	)

	err := cat.categoryRepository.DeleteCategory(uuid.MustParse(c.Param("id")))
	if err != nil {
		log.Error("Happened Error when delete category. Error: ", err)
		pkg.PanicException(constant.UnkonwnError)
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), ""))
}

func CategoryServiceInit(categoryRepository repositories.CategoryRepository) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		categoryRepository: categoryRepository,
	}
}
