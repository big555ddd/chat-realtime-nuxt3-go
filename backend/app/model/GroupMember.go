package model

import "github.com/uptrace/bun"

type GroupMember struct {
	bun.BaseModel `bun:"table:group_members"`

	GroupID  string `bun:"group_id" json:"group_id"`
	UserID   string `bun:"user_id" json:"user_id"`
	IsActive bool   `bun:"is_active" json:"is_active"`
	TimeRead int64  `bun:"time_read" json:"time_read"`
}
