package models

import (
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (model *GormModel) BeforeCreate(tx *gorm.DB) (err error) {
	model.CreatedAt = time.Now().Truncate(24 * time.Hour)
	return nil
}

func (model *GormModel) BeforeUpdate(tx *gorm.DB) (err error) {
	model.UpdatedAt = time.Now().Truncate(24 * time.Hour)
	return nil
}
