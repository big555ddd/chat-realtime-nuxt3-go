package routes

import (
	"app/app/controller"

	"github.com/gin-gonic/gin"
)

func SocketGroup(router *gin.RouterGroup) {
	ctl := controller.New()
	// md := middleware.AuthMiddleware()
	go ctl.GroupChatCtl.StartBroadcast()

	socket := router.Group("")
	{
		socket.GET("/ws/:id", ctl.GroupChatCtl.WebSocketHandler)
	}
}
