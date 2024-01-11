package models

type Model struct {
	ID        uint `gorm:"primarykey" json:"-"`
	CreatedAt int64
	UpdatedAt int64
}
