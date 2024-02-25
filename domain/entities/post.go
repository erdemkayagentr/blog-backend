package entities

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id            uuid.UUID `column:"id";json:"id";primaryKey:"true";default:"uuid_generate_v4()"`
	AuthorId      uuid.UUID `column:"author_id";json:"author_id"`
	CategoryId    uuid.UUID `column:"category_id";json:"category_id"`
	Title         string    `column:"title";json:"title"`
	SeoTitle      string    `column:"seo_title";json:"seo_title"`
	ImageUrl      string    `column:"image_url";json:"image_url"`
	Content       string    `column:"content";json:"content"`
	IsPublished   bool      `column:"is_published";json:"is_published"`
	PublishedDate time.Time `column:"published_date";json:"published_date"`
	CreatedAt     time.Time `column:"created_at";json:"created_at";default:"now()"`
}
