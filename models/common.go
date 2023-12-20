package models

import (
	"time"
)

type Model struct {
	ID        uint `gorm:"primarykey" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
