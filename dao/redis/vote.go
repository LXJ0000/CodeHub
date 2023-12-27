package redis

import (
	"bluebell/models"
	"bluebell/pkg/logger"
	"bluebell/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"math"
	"strconv"
	"time"
)

const (
	OneWeekInSecond = 7 * 24 * 3600
	ScorePreVote    = 432 // 每票的分值
)

/*
	基于用户投票的相关算法 https://ruanyifeng.com/blog/algorithm/
	投一票就加432分  86422/200 -> 200个赞则给帖子续一天 - 来自《Redis实战》

	- direction = 1
		- 没点过0      -> 更新分数和投票纪录          差值的绝对值为1  + 432
		- 点过反对-1    -> 更新分数和投票纪录          差值的绝对值为2   + 432 * 2
	- direction = 0
		- 取消反对-1    -> 更新分数和投票纪录           差值的绝对值为1  + 432

		- 取消赞成1    -> 更新分数和投票纪录           差值的绝对值为1  - 432
	- direction = -1
		- 没点过0      -> 更新分数和投票纪录          差值的绝对值为1  - 432
		- 点过赞成1    -> 更新分数和投票纪录          差值的绝对值为2  - 432 * 2

	投票的限制: 每个帖子自发表之日起一个星期内允许用户投票，超过一个星期不允许再投票
	1. 到期之后将Redis中保存的赞成票数以及反对票数存储到Mysql中
	2. 到期之后删除redis.KeyPostVotedPrefix
*/

func VoteForPost(c *gin.Context, userRID int64, req *models.VoteReq) {
	userID := strconv.FormatInt(userRID, 10)
	postID := strconv.FormatInt(req.PostID, 10)
	//	1. 判断是否超过帖子投票时间
	postTime := client.ZScore(KeyPostTimeZSet, postID).Val()
	if time.Now().Unix()-int64(postTime) > OneWeekInSecond {
		types.ResponseError(c, types.CodeVoteTimeExpire)
		return
	}

	//	2. 判断是否投过票
	key := KeyPostVotedPrefix + postID
	currDirection := client.ZScore(key, userID).Val() // 获取当前投票状态

	var scoreDirection float64
	if float64(req.Direction) > currDirection {
		scoreDirection = 1 // 加分
	} else {
		scoreDirection = -1 // 减分
	}
	diffAbs := math.Abs(currDirection - float64(req.Direction)) // 更新的倍数

	//开启事务 2 + 3
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(KeyPostScoreZSet, ScorePreVote*diffAbs*scoreDirection, postID)

	//	3. 记录用户为帖子投票的数据
	if req.Direction == 0 {
		//	删除投票记录
		pipeline.ZRem(KeyPostVotedPrefix+postID, userID)
	} else {
		pipeline.ZAdd(KeyPostVotedPrefix+postID, redis.Z{
			Score:  float64(req.Direction), // 赞成 or 反对 or 取消
			Member: userID,
		})
	}
	if _, err := pipeline.Exec(); err != nil {
		logger.Log.Error("事务回滚")
		types.ResponseError(c, types.CodeServerBusy)
		return
	}
	types.ResponseSuccess(c)
}

func CreatePostWithTime(postRID int64) error {
	postID := strconv.FormatInt(postRID, 10)
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
	_, err := pipeline.Exec()
	return err
}
