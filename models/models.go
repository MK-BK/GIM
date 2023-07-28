package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Models struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (m *Models) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New().String()
	return nil
}
