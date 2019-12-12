package utils

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
)

// We'll need to define an Upgrader this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WebsocketActor interface {
	Name() string
	Connection() *websocket.Conn
}

type wsClient struct {
	name string
	conn *websocket.Conn
}

func (c *wsClient) Name() string {
	return c.name
}

func (c *wsClient) Connection() *websocket.Conn {
	return c.conn
}

func (c *wsClient) dial(addr string) {
	u := url.URL{Scheme: "ws", Host: addr, Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	c.conn = conn
}

func NewWebsocketClient(serverAddr string) *wsClient {
	c := wsClient{
		name: "Websocket Client",
	}
	c.dial(serverAddr)
	return &c
}

type wsServer struct {
	name string
	conn *websocket.Conn
}

func (s *wsServer) Name() string {
	return s.name
}

func (s *wsServer) UpgradeConnection(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	s.conn = ws

	log.Println("Server started. Client connected.")
}

func (s *wsServer) Connection() *websocket.Conn {
	return s.conn
}

func NewWebsocketServer() *wsServer {
	s := wsServer{
		name: "Websocket Server",
	}
	return &s
}
