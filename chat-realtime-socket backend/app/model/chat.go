package model

import "github.com/uptrace/bun"

type Chat struct {
	bun.BaseModel `bun:"chats"`
	ID            string `bun:",default:gen_random_uuid(),pk" json:"id"`
	UserMapId     string `bun:"user_map_id" json:"user_map_id"`
	Sender        string `bun:"sender" json:"sender"`
	Recipient     string `bun:"recipient" json:"recipient"`
	Type          string `bun:"type" json:"type"`
	Detail        string `bun:"detail" json:"detail"`
	IsRead        bool   `bun:"is_read" json:"is_read"`
	CreateUpdateUnixTimestamp
}
