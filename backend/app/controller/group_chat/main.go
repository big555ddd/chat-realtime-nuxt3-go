package groupChat

import (
	"app/app/request"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/uptrace/bun"
)

type Controller struct {
	Name    string
	Service *Service
}

func NewController(db *bun.DB) *Controller {
	return &Controller{
		Name:    `group_chat-ctl`,
		Service: NewService(db),
	}
}

type Service struct {
	db *bun.DB
}

func NewService(db *bun.DB) *Service {
	return &Service{
		db: db,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow any origin; adjust for production
	},
}

type WebSocketManager struct {
	clients   map[*websocket.Conn]string
	broadcast chan request.SentMessage
	lock      sync.Mutex
}

var manager = WebSocketManager{
	clients:   make(map[*websocket.Conn]string),
	broadcast: make(chan request.SentMessage),
}
