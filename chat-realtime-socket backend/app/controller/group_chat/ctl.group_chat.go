package groupChat

import (
	"app/app/helper"
	"app/app/request"
	"app/app/response"
	"context"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) Create(c *gin.Context) {
	var req request.CreateGroup
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := ctl.Service.Create(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, data)
}

func (ctl *Controller) List(c *gin.Context) {
	req := request.GroupQuery{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if req.OrderBy == "" {
		req.OrderBy = "asc"
	}

	if req.SortBy == "" {
		req.SortBy = "created_at"
	}

	data, count, err := ctl.Service.List(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.SuccessWithPaginateAuto(c, data, req.Limit, req.Page, count)
}

func (ctl *Controller) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	data, err := ctl.Service.Get(c, id)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, data)
}

func (ctl *Controller) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}
	var req request.UpdateGroup
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := ctl.Service.Update(c, req, id)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, data)
}

func (ctl *Controller) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	err := ctl.Service.Delete(c, id)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, "Group Delete successfully")
}

func (ctl *Controller) ListMessages(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "id is required")
		return
	}

	req := request.GroupQuery{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if req.OrderBy == "" {
		req.OrderBy = "asc"
	}

	if req.SortBy == "" {
		req.SortBy = "created_at"
	}

	data, count, err := ctl.Service.ListMessages(c, id, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.SuccessWithPaginateAuto(c, data, req.Limit, req.Page, count)
}

func (ctl *Controller) SentMessage(c *gin.Context) {
	req := request.SentMessage{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	_, err := ctl.Service.SentMessage(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (ctl *Controller) WebSocketHandler(c *gin.Context) {
	// Upgrade the connection to WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}

	// Get the token from the request header
	user, err := helper.GetUserByToken(c)
	if err != nil {
		log.Println("Error fetching user:", err)
		conn.Close()
		return
	}

	req := request.GroupQuery{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if req.OrderBy == "" {
		req.OrderBy = "asc"
	}

	if req.SortBy == "" {
		req.SortBy = "created_at"
	}

	// Register the WebSocket connection with the user
	manager.lock.Lock()
	manager.clients[conn] = user
	manager.lock.Unlock()

	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, errors.New("id is required"))
		return
	}

	messages, _, err := ctl.Service.ListMessages(c, id, req)
	if err != nil {
		log.Println("Error fetching messages:", err)
		conn.Close()
		return
	}
	log.Println(messages)

	// Send the historical messages to the client over WebSocket
	for _, msg := range messages {
		err := conn.WriteJSON(msg) // Send message as JSON
		if err != nil {
			log.Println("Error sending message to client:", err)
			break
		}
	}

	// Listen for incoming WebSocket messages
	defer func() {
		manager.lock.Lock()
		delete(manager.clients, conn)
		manager.lock.Unlock()
		conn.Close()
	}()

	// Continuously read messages from the WebSocket connection
	for {
		var msg request.SentMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("WebSocket Read Error:", err)
			break
		}

		// Broadcast the received message to other clients
		manager.broadcast <- msg
	}
}

func (ctl *Controller) StartBroadcast() {
	go func() {
		for {
			msg := <-manager.broadcast // Get the message from the broadcast channel

			// Save the message to the database
			chat := request.SentMessage{
				Sender:  msg.Sender,
				GroupID: msg.GroupID,
				Detail:  msg.Detail,
				Type:    msg.Type,
			}

			log.Println("Broadcasting message:", chat)
			ctx := context.Background()
			createdChat, err := ctl.Service.SentMessage(ctx, chat)
			if err != nil {
				log.Println("Error saving message:", err)
				continue
			}

			// Include the created_at timestamp in the message
			createdChatDetail := map[string]interface{}{
				"sender":     createdChat.UserID,
				"group_id":   createdChat.GroupID,
				"detail":     createdChat.Detail,
				"type":       createdChat.Type,
				"created_at": createdChat.CreatedAt,
			}

			// Broadcast the message with created_at to all connected clients
			manager.lock.Lock()
			for conn := range manager.clients {
				err := conn.WriteJSON(createdChatDetail)
				if err != nil {
					log.Println("WebSocket Write Error:", err)
					conn.Close()
					delete(manager.clients, conn)
				}
			}
			manager.lock.Unlock()
		}
	}()
}
