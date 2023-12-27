package redis

const (
	KeyPrefix          = "bluebell:"
	KeyPostTimeZSet    = "post:time"
	KeyPostScoreZSet   = "post:score"
	KeyPostVotedPrefix = "post:voted:" // 参数 post_id
)
