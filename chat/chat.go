package chat

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

type ChatResponse struct {
    ChatMessage string `json:"chat_message"`
}

type HttpClient struct {
    w http.ResponseWriter
    r *http.Request
}

type Server struct {
    conns map[*websocket.Conn]bool
}

func NewServer() *Server {
    return &Server{
        conns: make(map[*websocket.Conn]bool),
    }
}

func (s *Server) handleChat(ws *websocket.Conn) {
    fmt.Println("new incomming connection from client:", ws.RemoteAddr())

    s.conns[ws] = true
    s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
    for {
        msg := ChatResponse{}
        err := websocket.JSON.Receive(ws, &msg)
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Println("read error:", err)
            continue
        }
        s.broadcast([]byte(msg.ChatMessage))
        fmt.Println(string(msg.ChatMessage))
    }
}

func (s *Server) broadcast(msg []byte) {
    for ws := range s.conns {
        go func(ws *websocket.Conn) {
            err := websocket.Message.Send(ws, `
                <div hx-swap-oob="beforeend:#replaceMe">
                    <p>` + string(msg) + `</p>
                </div>
                `)
            if err != nil {
                fmt.Println(err)
            }
        }(ws)
    }
}


func StartServer() {
    server := NewServer()
    http.Handle("/chat/get-messages", websocket.Handler(server.handleChat))
}

