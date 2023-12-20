package models

//Community

type Community struct {
	CommunityID   int64  `json:"community_id"`
	CommunityName string `json:"community_name"`
}

type CommunityInfo struct {
	Community
	Introduction string `json:"introduction,omitempty"`
	CreatedAt    int64  `json:"created_at,omitempty"`
	UpdatedAt    int64  `json:"updated_at,omitempty"`
}
