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

type PostCreateReq struct {
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	CommunityID int64  `json:"community_id" binding:"required"`
}

type PostListReq struct {
	Page
	CommunityID int64 `json:"community_id"`
}
