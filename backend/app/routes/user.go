package routes

import (
	"app/app/controller"
	"app/app/middleware"

	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup) {
	// Get the *bun.DB instance from config
	ctl := controller.New() // Pass the *bun.DB to the controller
	md := middleware.AuthMiddleware()
	user := router.Group("")
	{
		user.GET("/list", md, ctl.UserCtl.List)
		user.GET("/:id", md, ctl.UserCtl.Get)
		user.PATCH("/:id", md, ctl.UserCtl.Update)
		user.DELETE("/:id", md, ctl.UserCtl.Delete)
		user.GET("/listfriend/:id", md, ctl.UserCtl.ListFriend)
	}
}
