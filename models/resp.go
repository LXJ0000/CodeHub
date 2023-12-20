package models

import "time"

//Community

type Community struct {
	CommunityID   int64  `json:"community_id"`
	CommunityName string `json:"community_name"`
}

type CommunityInfo struct {
	Community
	Introduction string    `json:"introduction,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
