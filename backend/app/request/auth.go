package request

type ResetPassword struct {
	RedirectURL string `json:"redirect_url" binding:"required"`
	Email       string `json:"email" binding:"required"`
}

type ChangePasswordReset struct {
	Password string `json:"password"`
	Token    string `json:"token"`
}
