package redis

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

func (PostDao) GetPostIDInorder(req *models.PostListProReq) (ids []string, err error) {
	key := KeyPostScoreZSet
	if req.Order == models.OrderTime {
		key = KeyPostTimeZSet
	}
	//确定查询索引起始点
	start := (req.Page - 1) * req.Size
	end := (req.Page-1)*req.Size + req.Size - 1
	//按照 从大到小 查询 指定数量
	ids, err = client.ZRevRange(key, start, end).Result()
	return
}
