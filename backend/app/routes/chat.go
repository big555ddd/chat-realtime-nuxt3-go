package routes

import (
	"app/app/controller"
	"app/app/middleware"

	"github.com/gin-gonic/gin"
)

func Chat(router *gin.RouterGroup) {
	// Get the *bun.DB instance from config
	ctl := controller.New() // Pass the *bun.DB to the controller
	md := middleware.AuthMiddleware()
	chat := router.Group("")
	{
		chat.POST("/create", md, ctl.ChatCtl.Create)
		chat.GET("/:id", md, ctl.ChatCtl.GetMessage)
		chat.POST("/add-friend", md, ctl.ChatCtl.Addfriend)
		chat.POST("/remove-friend", md, ctl.ChatCtl.Removefriend)
		chat.GET("/unread-message/:id", md, ctl.ChatCtl.GetUnreadMessage)
		chat.POST("/read-message/:id", md, ctl.ChatCtl.ReadMessage)
	}
}
