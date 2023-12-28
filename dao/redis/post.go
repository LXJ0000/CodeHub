package redis

import (
	"bluebell/models"
	"github.com/go-redis/redis"
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

func (PostDao) GetPostVoteScore(ids []string) (scores []int64, err error) {
	//for _, id := range ids {
	//	key := KeyPostVotedPrefix + id
	//	//查找key中 投赞成票的数量
	//	scores = append(scores, client.ZCount(key, "1", "1").Val())
	//}
	//优化
	//使用pipeline一次发送多条命令 减少RTT
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := KeyPostVotedPrefix + id
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return
	}
	for _, cmder := range cmders {
		val := cmder.(*redis.IntCmd).Val()
		scores = append(scores, val)
	}
	return
}
