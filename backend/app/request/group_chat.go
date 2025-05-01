package request

type CreateGroup struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	UserID      []string `json:"user_id"`
}

type UpdateGroup struct {
	CreateGroup
	IsActive bool `json:"is_active"`
}

type AddOrRemoveUser struct {
	GroupID string   `json:"group_id"`
	UserID  []string `json:"user_id"`
}

type GroupQuery struct {
	Page     int    `form:"page"`
	Limit    int    `form:"limit"`
	SortBy   string `form:"sort_by"`
	OrderBy  string `form:"order_by"`
	Search   string `form:"search"`
	SearchBy string `form:"search_by"`
	UserID   string `form:"user_id"`
	ID       string
}

type SentMessage struct {
	Sender  string `bun:"sender" json:"sender"`
	GroupID string `bun:"group_id" json:"group_id"`
	Type    string `bun:"type" json:"type"`
	Detail  string `bun:"detail" json:"detail"`
}
