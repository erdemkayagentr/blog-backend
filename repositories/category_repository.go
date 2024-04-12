package repositories

import (
	"erdemkayacomtr/domain/entities"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategories() ([]entities.Category, error)
	GetCategoryById(id uuid.UUID) (entities.Category, error)
	CreateCategory(category *entities.Category) (entities.Category, error)
	UpdateCategory(category *entities.Category) (entities.Category, error)
	DeleteCategory(id uuid.UUID) error
}
type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func (c CategoryRepositoryImpl) GetCategories() ([]entities.Category, error) {
	var categories []entities.Category

	var err = c.db.Find(&categories).Error
	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}
	return categories, nil
}

func (c CategoryRepositoryImpl) GetCategoryById(id uuid.UUID) (entities.Category, error) {
	var category entities.Category
	var err = c.db.Where("id = ?", id).Find(&category, id).Error
	if err != nil {
		return entities.Category{}, err
	}
	return category, nil
}

func (c CategoryRepositoryImpl) CreateCategory(category *entities.Category) (entities.Category, error) {
	var err = c.db.Create(&category).Error
	if err != nil {
		return entities.Category{}, err
	}
	return *category, nil
}

func (c CategoryRepositoryImpl) UpdateCategory(category *entities.Category) (entities.Category, error) {
	var err = c.db.Save(&category).Error
	if err != nil {
		return entities.Category{}, err
	}
	return *category, nil
}

func (c CategoryRepositoryImpl) DeleteCategory(id uuid.UUID) error {
	var err = c.db.Delete(&entities.Category{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func CategoryRepositoryInit(db *gorm.DB) *CategoryRepositoryImpl {
	db.AutoMigrate(&entities.Category{})
	return &CategoryRepositoryImpl{
		db: db,
	}
}
