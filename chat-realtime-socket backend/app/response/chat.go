package response

type ChatUnreadResponse struct {
	Unread int    `json:"unread"`
	Type   string `json:"type"`
	Detail string `json:"detail"`
}
