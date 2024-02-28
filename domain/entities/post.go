package entities

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID            uuid.UUID `gorm:"column:id;type:uuid;primaryKey:true;default:uuid_generate_v4();not null:true;" json:"id"`
	AuthorId      uuid.UUID `gorm:"column:author_id;type:uuid;not null:true;" json:"author_id"`
	CategoryId    uuid.UUID `gorm:"column:category_id;type:uuid;not null:true;" json:"category_id"`
	Title         string    `gorm:"column:title;type:varchar(255);not null:true;" json:"title"`
	SeoTitle      string    `gorm:"column:seo_title;type:varchar(255);not null:true;" json:"seo_title"`
	ImageUrl      string    `gorm:"column:image_url;type:varchar(255);not null:true;" json:"image_url"`
	Content       string    `gorm:"column:content;type:text;not null:true;" json:"content"`
	IsPublished   bool      `gorm:"column:is_published;type:boolean;not null:true;" json:"is_published"`
	PublishedDate time.Time `gorm:"column:published_date;type:timestamp;not null:true;" json:"published_date"`
	Category      `gorm:"foreignKey:CategoryId;references:ID" json:"category"`
	BaseModel
}
