package mysql

import (
	"bluebell/models"
	"strings"
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

// Create 创建帖子
func (p *PostDao) Create(post *models.PostModel) error {
	return db.Create(&post).Error
}

// GetList 根据给定限制条件查询帖子列表
func (p *PostDao) GetList(condition map[string]interface{}, page *models.Page) (list []*models.PostResp, err error) {
	err = db.Model(&models.PostModel{}).Where(condition).Offset(page.Size * (page.Num - 1)).Limit(page.Size).Find(&list).Error
	return
}

// GetInfo 根据当个ID查询帖子详细信息
func (p *PostDao) GetInfo(pid int64) (info *models.PostResp, err error) {
	err = db.Model(&models.PostModel{}).Where("post_id=?", pid).First(&info).Error
	return
}

// GetCountByCondition 获取符合条件的帖子数目
func (p *PostDao) GetCountByCondition(condition map[string]interface{}) (total int64, err error) {
	err = db.Model(&models.PostModel{}).Where(condition).Count(&total).Error
	return
}

// GetPostListWithIDList 根据给定ID列表查询帖子详细信息
func (p *PostDao) GetPostListWithIDList(ids []string) (list []*models.PostResp, err error) {
	//err = db.Model(&models.PostModel{}).Where("post_id in ?", ids).Order("created_at DESC").Find(&list).Error
	agr := strings.Join(ids, ",")
	err = db.Raw("select post_id,title,content,author_id,community_id,created_at from post where post_id in ? order by FIND_IN_SET(post_id, ?)", ids, agr).Scan(&list).Error
	return

}
