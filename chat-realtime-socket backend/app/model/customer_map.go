package model

import "github.com/uptrace/bun"

type UserMap struct {
	bun.BaseModel `bun:"user_maps"`
	ID            string `bun:",default:gen_random_uuid(),pk" json:"id"`
	UserID        string `bun:"user_id" json:"user_id"`
	RefID         string `bun:"ref_id" json:"ref_id"`
	CreateUpdateUnixTimestamp
}
