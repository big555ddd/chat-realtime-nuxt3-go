package chat

import (
	"app/app/controller/user"
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

func NewController(db *bun.DB, userservice *user.Controller) *Controller {
	return &Controller{
		Name:    `chat-ctl`,
		Service: NewService(db, userservice),
	}
}

type Service struct {
	db          *bun.DB
	userservice *user.Controller
}

func NewService(db *bun.DB, userservice *user.Controller) *Service {
	return &Service{
		db:          db,
		userservice: userservice,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow any origin; adjust for production
	},
}

type Client struct {
	conn      *websocket.Conn
	userID    string
	contactID string
}

var manager = struct {
	clients   map[*Client]bool
	broadcast chan request.CreateChat
	lock      sync.Mutex
}{
	clients:   make(map[*Client]bool),
	broadcast: make(chan request.CreateChat),
}
