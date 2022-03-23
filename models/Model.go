package models

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:varchar(36)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by" gorm:"type:varchar(32)"`
	UpdatedBy string    `json:"updated_by" gorm:"type:varchar(32)"`
	Deleted   int8      `json:"deleted" gorm:"default:0"`
}
