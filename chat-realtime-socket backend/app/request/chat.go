package request

type CreateChat struct {
	Sender    string `bun:"sender" json:"sender"`
	Recipient string `bun:"recipient" json:"recipient"`
	Type      string `bun:"type" json:"type"`
	Detail    string `bun:"detail" json:"detail"`
}

type Addfriend struct {
	UserID string `bun:"user_id" json:"user_id"`
}

type GetUnreadMessage struct {
	Id string `uri:"id"`
}
