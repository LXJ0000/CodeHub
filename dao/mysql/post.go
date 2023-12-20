package mysql

import (
	"bluebell/models"
	"sync"
)

type PostDao struct {
}

var (
	postDao  *PostDao
	postOnce sync.Once
)

func NewPostDao() *PostDao {
	postOnce.Do(func() {
		postDao = &PostDao{}
	})
	return postDao
}

func (p *PostDao) Create(post *models.PostModel) error {
	return db.Create(&post).Error
}
