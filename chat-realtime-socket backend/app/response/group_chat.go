package response

type GroupChatResponse struct {
	ID                     string               `bun:"id" json:"id"`
	Name                   string               `bun:"name" json:"name"`
	Description            string               `bun:"description" json:"description"`
	IsActive               bool                 `bun:"is_active" json:"is_active"`
	CreatedAt              int64                `bun:"created_at" json:"created_at"`
	LastMessageTime        int64                `bun:"last_message_time" json:"last_message_time"`
	LastMessageSender      string               `bun:"last_message_sender" json:"last_message_sender"`
	LastMessageType        string               `bun:"last_message_type" json:"last_message_type"`
	LastMessageDetail      string               `bun:"last_message_detail" json:"last_message_detail"`
	LastMessageDisplayName string               `bun:"last_message_display_name" json:"last_message_display_name"`
	LastMessageFirstName   string               `bun:"last_message_first_name" json:"last_message_first_name"`
	Member                 []MemberChatResponse `bun:"members" json:"members"`
}
type MemberInfo struct {
	GroupID     string
	UserID      string
	TimeRead    int64
	IsActive    bool
	FirstName   string
	LastName    string
	Username    string
	DisplayName string
}


type MemberChatResponse struct {
	ID           string `bun:"id" json:"id"`
	Username     string `bun:"username" json:"username"`
	Firstname    string `bun:"first_name" json:"first_name"`
	DisplayName  string `bun:"display_name" json:"display_name"`
	Lastname     string `bun:"last_name" json:"last_name"`
	IsActive     bool   `bun:"is_active" json:"is_active"`
	TimeRead     int64  `bun:"time_read" json:"time_read"`
	MessageCount int64  `bun:"message_count" json:"message_count"`
}

type GroupChatMap struct {
	UserID      string `bun:"user_id" json:"user_id"`
	Firstname   string `bun:"first_name" json:"first_name"`
	DisplayName string `bun:"display_name" json:"display_name"`
	Lastname    string `bun:"last_name" json:"last_name"`
	IsActive    bool   `bun:"is_active" json:"is_active"`
	Type        string `bun:"type" json:"type"`
	Detail      string `bun:"detail" json:"detail"`
	CreatedAt   int64  `bun:"created_at" json:"created_at"`
}


