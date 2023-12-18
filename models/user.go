package models

type UserModel struct {
	Model
	UserID   int64  `gorm:"not null"`
	UserName string `gorm:"not null"`
	NickName string
	Password string `gorm:"not null"`
	Email    string
	Gender   bool `gorm:"default:true"` // true man
}
