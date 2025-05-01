package routes

import (
	"app/app/controller"

	"github.com/gin-gonic/gin"
)

func Socket(router *gin.RouterGroup) {
	ctl := controller.New()
	// md := middleware.AuthMiddleware()
	go ctl.ChatCtl.StartBroadcast()

	socket := router.Group("")
	{
		socket.GET("/ws/:id", ctl.ChatCtl.WebSocketHandler)
	}
}
