package redis

const (
	KeyPrefix          = "bluebell:"
	KeyPostTimeZSet    = KeyPrefix + "post:time"
	KeyPostScoreZSet   = KeyPrefix + "post:score"
	KeyPostVotedPrefix = KeyPrefix + "post:voted:" // 参数 post_id

	KeyCommunitySetPrefix = KeyPrefix + "community:" // 参数 community_id
)
