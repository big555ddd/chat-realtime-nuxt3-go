package user

import (
	"app/app/request"
	"app/app/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) Create(c *gin.Context) {
	var req request.CreateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctl.Service.Create(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, user)
}

func (ctl *Controller) List(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		response.BadRequest(c, "limit is invalid")
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		response.BadRequest(c, "page is invalid")
		return
	}
	search := c.DefaultQuery("search", "")
	roleID := c.DefaultQuery("role_id", "")
	status := c.DefaultQuery("status", "")
	plan_type := c.DefaultQuery("plan_type", "")
	users, count, err := ctl.Service.List(c.Request.Context(), limit, page, search, roleID, status, plan_type)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.SuccessWithPaginateAuto(c, users, limit, page, count)
}

func (ctl *Controller) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	user, err := ctl.Service.Get(c.Request.Context(), id)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, user)
}

func (ctl *Controller) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}
	var req request.UpdateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := ctl.Service.Update(c.Request.Context(), req, id)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, nil)
}

func (ctl *Controller) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	err := ctl.Service.Delete(c.Request.Context(), id)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, nil)
}

func (ctl *Controller) ListFriend(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		response.BadRequest(c, "limit is invalid")
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		response.BadRequest(c, "page is invalid")
		return
	}
	search := c.DefaultQuery("search", "")

	users, count, err := ctl.Service.ListFriend(c.Request.Context(), limit, page, id, search)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.SuccessWithPaginateAuto(c, users, limit, page, count)
}
