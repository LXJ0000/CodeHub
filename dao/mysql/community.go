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
