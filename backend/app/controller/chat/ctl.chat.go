package chat

import (
	"app/app/helper"
	"app/app/request"
	"app/app/response"
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) Create(c *gin.Context) {

	user, err := helper.GetUserByToken(c)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	req := request.CreateChat{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req.Sender = user

	data, err := ctl.Service.Create(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, data)
}

func (ctl *Controller) GetMessage(c *gin.Context) {

	user, err := helper.GetUserByToken(c)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, errors.New("id is required"))
		return
	}

	data, count, err := ctl.Service.GetMessage(c.Request.Context(), limit, page, user, id)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.SuccessWithPaginateAuto(c, data, limit, page, count)
}

func (ctl *Controller) ReadMessage(c *gin.Context) {

	user, err := helper.GetUserByToken(c)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	query := new(request.GetUnreadMessage)
	if err := c.BindUri(query); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err = ctl.Service.ReadMessage(c, user, query.Id)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, nil)
}

func (ctl *Controller) GetUnreadMessage(c *gin.Context) {

	user, err := helper.GetUserByToken(c)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	query := new(request.GetUnreadMessage)
	if err := c.BindUri(query); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := ctl.Service.GetUnreadMessage(c, user, query.Id)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, data)
}

func (ctl *Controller) Addfriend(c *gin.Context) {

	user, err := helper.GetUserByToken(c)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	req := request.Addfriend{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	data, err := ctl.Service.Addfriend(c.Request.Context(), req, user)

	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, data)

}

func (ctl *Controller) Removefriend(c *gin.Context) {

	user, err := helper.GetUserByToken(c)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	req := request.Addfriend{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	data, err := ctl.Service.Removefriend(c.Request.Context(), req, user)

	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, data)

}

func (ctl *Controller) WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}

	user, err := helper.GetUserByToken(c)
	if err != nil {
		log.Println("Error fetching user:", err)
		conn.Close()
		return
	}

	contactID := c.Param("id")
	if contactID == "" {
		response.BadRequest(c, errors.New("id is required"))
		conn.Close()
		return
	}

	client := &Client{
		conn:      conn,
		userID:    user,      // เก็บ userID ของผู้ใช้
		contactID: contactID, // เก็บ contactID ที่เชื่อมต่อ
	}

	manager.lock.Lock()
	manager.clients[client] = true // เพิ่ม client ใหม่เข้าไปใน manager
	manager.lock.Unlock()

	// ส่งข้อความเก่าให้ client
	messages, _, err := ctl.Service.GetMessage(c, 20, 1, user, contactID)
	if err != nil {
		log.Println("Error fetching messages:", err)
		conn.Close()
		return
	}

	for _, msg := range messages {
		err := conn.WriteJSON(msg)
		if err != nil {
			log.Println("Error sending message to client:", err)
			break
		}
	}

	defer func() {
		manager.lock.Lock()
		delete(manager.clients, client) // ลบ client ออกเมื่อเชื่อมต่อปิด
		manager.lock.Unlock()
		conn.Close()
	}()

	for {
		var msg request.CreateChat
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("WebSocket Read Error:", err)
			break
		}

		// Broadcast ข้อความไปยัง client ที่เกี่ยวข้อง
		manager.broadcast <- msg
	}
}

func (ctl *Controller) StartBroadcast() {
	go func() {
		for {
			msg := <-manager.broadcast

			// Save the message to the database
			chat := request.CreateChat{
				Sender:    msg.Sender,
				Recipient: msg.Recipient,
				Detail:    msg.Detail,
				Type:      msg.Type,
			}

			ctx := context.Background()
			createdChat, err := ctl.Service.Create(ctx, chat)
			if err != nil {
				log.Println("Error saving message:", err)
				continue
			}

			err = ctl.Service.ReadMessage(ctx, createdChat.Sender, createdChat.Recipient)
			if err != nil {
				log.Println("Error marking message as read:", err)
				continue
			}

			total, err := ctl.Service.GetUnreadMessage(ctx, createdChat.Recipient, createdChat.Sender)
			if err != nil {
				log.Println("Error fetching unread message:", err)
				continue
			}

			type chat5 struct {
				Sender    string `json:"sender"`
				Recipient string `json:"recipient"`
				Detail    string `json:"detail"`
				Type      string `json:"type"`
				Total     int    `json:"total"`
				CreatedAt int64  `json:"created_at"`
			}

			// createdChatDetail := map[string]interface{}{
			// 	"sender":     createdChat.Sender,
			// 	"recipient":  createdChat.Recipient,
			// 	"detail":     createdChat.Detail,
			// 	"type":       createdChat.Type,
			// 	"total":      total,
			// 	"created_at": createdChat.CreatedAt,
			// }
			createdChatDetail := chat5{
				Sender:    createdChat.Sender,
				Recipient: createdChat.Recipient,
				Detail:    createdChat.Detail,
				Type:      createdChat.Type,
				Total:     total,
				CreatedAt: createdChat.CreatedAt,
			}

			manager.lock.Lock()
			for client := range manager.clients {

				// ส่งข้อความเฉพาะให้กับ client ที่เกี่ยวข้อง
				if client.userID == msg.Sender || client.userID == msg.Recipient {
					err := client.conn.WriteJSON(createdChatDetail)
					if err != nil {
						log.Println("WebSocket Write Error:", err)
						client.conn.Close()
						delete(manager.clients, client)
					}
				}
			}
			manager.lock.Unlock()
		}
	}()
}
