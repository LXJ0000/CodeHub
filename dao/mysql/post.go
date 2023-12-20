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

func (p *PostDao) GetList() (list []*models.PostResp, err error) {
	err = db.Model(&models.PostModel{}).Find(&list).Error
	return
}

func (p *PostDao) GetInfo(pid int64) (info *models.PostResp, err error) {
	err = db.Model(&models.PostModel{}).Where("post_id=?", pid).First(&info).Error
	return
}
