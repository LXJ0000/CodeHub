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

func (p *PostDao) GetList(condition map[string]interface{}, page *models.Page) (list []*models.PostResp, err error) {
	err = db.Model(&models.PostModel{}).Where(condition).Offset(page.Size * (page.Num - 1)).Limit(page.Size).Find(&list).Error
	return
}

func (p *PostDao) GetInfo(pid int64) (info *models.PostResp, err error) {
	err = db.Model(&models.PostModel{}).Where("post_id=?", pid).First(&info).Error
	return
}

func (p *PostDao) GetCountByCondition(condition map[string]interface{}) (total int64, err error) {
	err = db.Model(&models.PostModel{}).Where(condition).Count(&total).Error
	return
}
