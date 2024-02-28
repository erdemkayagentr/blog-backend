package entities

import (
	"github.com/google/uuid"
)

type Category struct {
	ID    uuid.UUID `gorm:"column:id;type:uuid;primaryKey:true;default:uuid_generate_v4();not null:true;" json:"id"`
	Name  string    `gorm:"column:name;type:varchar(255);not null:true;" json:"name"`
	Posts []Post    `gorm:"foreignKey:CategoryId;references:ID" json:"-"`
	BaseModel
}
