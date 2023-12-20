package mysql

import (
	"bluebell/models"
	"sync"
)

type CommunityDao struct {
}

var (
	communityDao  *CommunityDao
	communityOnce sync.Once
)

func NewCommunityDao() *CommunityDao {
	communityOnce.Do(func() {
		communityDao = &CommunityDao{}
	})
	return communityDao
}

func (c *CommunityDao) GetList() (list []*models.Community, err error) {
	err = db.Model(&models.CommunityModel{}).Find(&list).Error
	return
}

func (c *CommunityDao) GetInfo(id int64) (info *models.CommunityInfo, err error) {
	err = db.Model(&models.CommunityModel{}).Where("community_id=?", id).First(&info).Error
	return
}
