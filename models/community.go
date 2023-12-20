package models

type CommunityModel struct {
	Model
	CommunityID   int64  `gorm:"not null"`
	CommunityName string `gorm:"not null"`
	Introduction  string `gorm:""`
}

func (CommunityModel) TableName() string {
	return `community`
}
