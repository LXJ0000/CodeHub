package redis

import (
	"bluebell/models"
	"github.com/go-redis/redis"
	"strconv"
	"sync"
	"time"
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

func getIdsFromKey(key string, page, size int64) ([]string, error) {
	//确定查询索引起始点
	start := (page - 1) * size
	end := (page-1)*size + size - 1
	//按照 从大到小 查询 指定数量
	return client.ZRevRange(key, start, end).Result()
}

// GetPostIDInorder 按要求按序获取post_id
func (PostDao) GetPostIDInorder(req *models.PostListProReq) (ids []string, err error) {
	var key string

	//1. 判断是否有community_id 字段
	if req.CommunityID != 0 {
		//存在 community_id 字段
		communityID := strconv.FormatInt(req.CommunityID, 10)

		//使用ZInterStore 把 SET community:community_id 与 ZSET post:score | post:time 生成一个新的ZSet，里面是符合要求的post_id

		//针对新的ZSet 按之前的逻辑取数据

		communityKey := KeyCommunitySetPrefix + communityID //community:community_id
		var orderKey string                                 //post:time or post:score
		if req.Order == "time" {
			orderKey = KeyPostTimeZSet
		} else {
			orderKey = KeyPostScoreZSet
		}

		//利用缓存key减少ZInterStore执行的次数 - 需要设置超时时间
		key = orderKey + ":" + communityID
		if client.Exists(key).Val() < 1 {
			//	不存在 则需要计算
			pipeline := client.TxPipeline()
			pipeline.ZInterStore(key, redis.ZStore{
				Weights:   nil,
				Aggregate: "MAX",
			}, communityKey, orderKey)
			pipeline.Expire(key, 60*time.Second) // 设置超时时间
			if _, err = pipeline.Exec(); err != nil {
				return
			}
		}
	} else {
		//不存在 community_id 字段 直接使用 post:time || post:score 查询
		key = KeyPostScoreZSet // post:score
		if req.Order == models.OrderTime {
			key = KeyPostTimeZSet // post:time
		}
	}
	return getIdsFromKey(key, req.Page, req.Size)

}

// GetPostVoteScore 获取帖子分数列表
func (PostDao) GetPostVoteScore(ids []string) (scores []int64, err error) {
	pipeline := client.TxPipeline()
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

// CreatePost 帖子部分信息加入Redis time、score、community_id
func (PostDao) CreatePost(postRID int64, communityRID int64) error {
	postID := strconv.FormatInt(postRID, 10)
	communityID := strconv.FormatInt(communityRID, 10)
	//开启事务
	pipeline := client.TxPipeline()
	//帖子时间
	pipeline.ZAdd(KeyPostTimeZSet, redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	//帖子分数
	pipeline.ZAdd(KeyPostScoreZSet, redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	//把帖子id加到社区set
	pipeline.SAdd(KeyCommunitySetPrefix+communityID, postID)
	_, err := pipeline.Exec()
	return err
}
