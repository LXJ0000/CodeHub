package models

//社区

type CommunityResp struct {
	CommunityID   int64  `json:"community_id"`
	CommunityName string `json:"community_name"`
}

type CommunityInfoResp struct {
	CommunityResp
	Introduction string `json:"introduction,omitempty"`
	CreatedAt    int64  `json:"created_at,omitempty"`
	UpdatedAt    int64  `json:"updated_at,omitempty"`
}
