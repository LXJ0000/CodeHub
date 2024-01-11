package models

//用户

type UserResp struct {
	UserID   int64  `json:"user_id,string"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
}

type UserInfoResp struct {
	*UserResp
	Email  string `json:"email"`
	Gender bool   `json:"gender"`
}

//社区

type CommunityResp struct {
	CommunityID   int64  `json:"community_id"`
	CommunityName string `json:"community_name"`
}

type CommunityInfoResp struct {
	*CommunityResp
	Introduction string `json:"introduction,omitempty"`
	//CreatedAt    int64  `json:"created_at,omitempty"`
	//UpdatedAt    int64  `json:"updated_at,omitempty"`
}

// 帖子

type PostResp struct {
	PostID      int64  `json:"post_id,string"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	AuthorID    int64  `json:"author_id,string"`
	CommunityID int64  `json:"community_id"`
	CreatedAt   int64  `json:"created_at"`
}

type PostInfoResp struct {
	AuthorName string `json:"author_name"`
	Score      int64  `json:"score"`
	*PostResp
	*CommunityInfoResp `json:"community"`
}
