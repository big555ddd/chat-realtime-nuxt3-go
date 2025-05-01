package helper

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type userID struct {
	AuthTyoe string `json:"auth_type"`
	ID       string `json:"id"`
	Data     struct {
		ID    string `json:"id"`
		Name  string `json:"username"`
		Email string `json:"email"`
	} `json:"data"`
	Nbf int64 `json:"nbf"`
	Exp int64 `json:"exp"`
}

func GetUserByToken(c *gin.Context) (string, error) {
	userData, _ := c.Get("claims")
	userBytes, _ := json.Marshal((userData))
	var usr userID
	err := json.Unmarshal(userBytes, &usr)

	return usr.ID, err
}
