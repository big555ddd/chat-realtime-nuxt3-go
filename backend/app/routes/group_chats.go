package routes

import (
	"app/app/controller"
	"app/app/middleware"

	"github.com/gin-gonic/gin"
)

func GroupChat(router *gin.RouterGroup) {
	// Get the *bun.DB instance from config
	ctl := controller.New() // Pass the *bun.DB to the controller
	md := middleware.AuthMiddleware()

	groupchat := router.Group("")
	{
		groupchat.POST("/create", md, ctl.GroupChatCtl.Create)
		groupchat.POST("/sent-message", md, ctl.GroupChatCtl.SentMessage)
		groupchat.PATCH("/:id", md, ctl.GroupChatCtl.Update)
		groupchat.GET("message/:id", md, ctl.GroupChatCtl.ListMessages)
		groupchat.GET("/list", md, ctl.GroupChatCtl.List)
		groupchat.GET("/:id", md, ctl.GroupChatCtl.Get)
	}
}
