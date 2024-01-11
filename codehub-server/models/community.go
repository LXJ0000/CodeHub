package models

type CommunityModel struct {
	Model
	CommunityID   int64  `gorm:"not null,unique"`
	CommunityName string `gorm:"not null,unique"`
	Introduction  string `gorm:""`
}

func (CommunityModel) TableName() string {
	return `community`
}
