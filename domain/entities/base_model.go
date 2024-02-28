package entities

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp;not null;default:now()" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp;null;default:CURRENT_TIMESTAMP" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;index" json:"-"`
}
