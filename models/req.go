package models

//Common

type Page struct {
	Size int `json:"size"`
	Num  int `json:"num"`
}

//用户

type UserRegisterReq struct {
	Username   string `json:"user_name" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type UserLoginReq struct {
	Username string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 帖子

const (
	OrderTime  = "time"
	OrderScore = "score"
)

type PostCreateReq struct {
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	CommunityID int64  `json:"community_id" binding:"required"`
}

type PostListReq struct {
	Page
	CommunityID int64 `json:"community_id"`
}

type VoteReq struct {
	PostID    int64 `json:"post_id,string" binding:"required"`
	Direction int8  `json:"direction,string" binding:"oneof=0 1 -1"` //赞成 1 or 反对 -1 or 取消投票 0
}

type PostListProReq struct {
	Page        int64  `form:"page" example:"1"`      // 页码 可以为空
	Size        int64  `form:"size" example:"5"`      // 每一页的数量 可以为空
	Order       string `form:"order" example:"score"` // 排序依据：order || time 可以为空
	CommunityID int64  `form:"community_id"`          // 是否按社区排序 可以为空
}
