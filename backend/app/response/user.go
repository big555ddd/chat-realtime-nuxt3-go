package response

type UserResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Firstname   string `json:"firstname"`
	DisplayName string `json:"display_name"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	RoleID      uint   `json:"role_id"`
	Status      string `json:"status"`
}

type UserListResponse struct {
	Users      []UserResponse `json:"users"`
	Pagination Pagination     `json:"pagination"`
}

type GetUserDetail struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Nickname  string `json:"nickname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	RoleID    uint   `json:"role_id"`
	Point     int64  `json:"points"`
}

type ListFriendResponse struct {
	UserResponse
	Type        string `json:"type"`
	Detail      string `json:"detail"`
	MessageTime int64  `json:"message_time"`
	Unread      int    `json:"unread"`
}
