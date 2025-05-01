package model

import (
	"github.com/uptrace/bun"
)

type GroupChat struct {
	bun.BaseModel `bun:"table:group_chats"`

	ID                string `bun:",default:gen_random_uuid(),pk" json:"id"`
	Name              string `bun:"name" json:"name"`
	Description       string `bun:"description" json:"description"`
	IsActive          bool   `bun:"is_active" json:"is_active"`
	LastMessageTime   int64  `bun:"last_message_time" json:"last_message_time"`
	LastMessageType   string `bun:"last_message_type" json:"last_message_type"`
	LastMessageDetail string `bun:"last_message_detail" json:"last_message_detail"`
	LastMessageSender string `bun:"last_message_sender" json:"last_message_sender"`
	CreateUpdateUnixTimestamp
	SoftDelete
}
