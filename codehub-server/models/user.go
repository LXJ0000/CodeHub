package models

type UserModel struct {
	Model
	UserID   int64  `gorm:"not null,unique"`
	UserName string `gorm:"not null,unique"`
	NickName string
	Password string `gorm:"not null"`
	Email    string
	Gender   bool `gorm:"default:true"` // true man
}

func (UserModel) TableName() string {
	return `user`
}
