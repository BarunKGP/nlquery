package core

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/BarunKGP/nlquery/internal"
	_ "github.com/joho/godotenv"

	"github.com/gorilla/websocket"
)

var logger = internal.CreateLogger()

const (
	writeWait      = 10 * time.Second
	pongWait       = 10 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSise = 1024
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		if r.URL.Host == os.Getenv("CLIENTSIDE_AUTH_REDIRECT_URL") {
			return true
		}
		return false
	},
}

type Client struct {
	Id   int64 `json:"id"`
	hub  *ChatHub
	conn *websocket.Conn
	send chan []byte
}

func NewClient(id int64, hub *ChatHub, conn *websocket.Conn, send chan []byte) *Client {
	return &Client{Id: id, hub: hub, conn: conn, send: send}
}

type llmClient interface {
	AskLlm([]byte, any) ([]byte, error)
	validateResponse([]byte) (bool, error)
}

type LlmChatClient struct {
	Client
	ModelUrl  string `json:"modelUrl,omitempty"`
	LlmParams any    `json:"modelParams,omitempty"` // TODO: This should be a struct controlling LLM params
}

// Builder pattern
// TODO: replace `llmParams` any type with final struct type
func (c *Client) AsLlm(modelUrl string, llmParams any) *LlmChatClient {
	return &LlmChatClient{Client: *c, ModelUrl: modelUrl, LlmParams: llmParams}
}

// TODO: Implement LlmChatClient as llmClient interface

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() error {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSise)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		logger.Debug(fmt.Sprintf("Pong received on client %d. Incrementing read deadline by %s", c.Id, pongWait))
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Debug(err.Error())
				return fmt.Errorf("Client closed unexpectedly!")
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- message
	}
	return nil
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	//? Besides writing messages, what do we need to do here?
	//* We need to make LLM calls, validate user input, prevent harmful/jailbreak responses and analytics...
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				logger.Error("Error sending ping: %v", err)
				return
			}
		}
	}
}

type ChatHub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func NewChatHub() *ChatHub {
	return &ChatHub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *ChatHub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			} else {
				logger.Warn("client not found: nothing to unregister", "clientId", client.Id)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
