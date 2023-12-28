package mysql

import (
	"bluebell/conf"
	"bluebell/models"
	"testing"
)

func TestPostDao_Create(t *testing.T) {
	conf.Init()
	Init()
	postDao := NewPostDao()
	post := &models.PostModel{
		Title:       "123",
		Content:     "123",
		PostID:      123,
		AuthorID:    123,
		CommunityID: 123,
		Status:      123,
	}
	err := postDao.Create(post)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("success")
}
