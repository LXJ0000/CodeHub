package models

type PostModel struct {
	Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`

	PostID      int64 `gorm:"not null,unique"`
	AuthorID    int64 `gorm:"not null"`
	CommunityID int64 `gorm:"not null"`

	Status uint8 `gorm:"not null,default:1"` // 审核、已审核
}

func (PostModel) TableName() string {
	return `post`
}
