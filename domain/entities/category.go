package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	// CreatedAt represents the timestamp when the entity was created.
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp;not null;default:now()" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp;null;default:CURRENT_TIMESTAMP" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;index" json:"-"`
}

type Category struct {
	ID    uuid.UUID `gorm:"column:id;type:uuid;primaryKey:true;default:uuid_generate_v4();not null:true;" json:"id"`
	Name  string    `gorm:"column:name;type:varchar(255);not null:true;" json:"name"`
	Posts []Post    `json:"-"`
	BaseModel
}
