package model

import "github.com/uptrace/bun"

type GroupChatMap struct {
	bun.BaseModel `bun:"table:group_chat_maps"`

	GroupID string `bun:"group_id" json:"group_id"`
	UserID  string `bun:"user_id" json:"user_id"`
	Type    string `bun:"type" json:"type"`
	Detail  string `bun:"detail" json:"detail"`
	CreateUpdateUnixTimestamp
}
