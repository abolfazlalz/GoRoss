package websocket

import (
	"github.com/googollee/go-socket.io"
	"goross/pkg/events"
	"log"
	"net/http"
)

type WebSocket struct {
	*events.AbstractEventListener[string]
	server *socketio.Server
}

var ws *WebSocket

func New() *WebSocket {
	ws = &WebSocket{
		server:                socketio.NewServer(nil),
		AbstractEventListener: events.NewAbstractEventListener[string](),
	}
	return ws
}

func Instance() *WebSocket {
	if ws == nil {
		return New()
	}
	return ws
}

func (ws *WebSocket) Server() *socketio.Server {
	return ws.server
}

func (ws *WebSocket) handleConnect(s socketio.Conn) error {
	s.SetContext("")
	log.Println("connected:", s.ID())
	return nil
}

func (ws *WebSocket) handleNotice(s socketio.Conn, msg string) {
	log.Println("notice:", msg)
	s.Emit("reply", "have "+msg)
}

func (ws *WebSocket) handleClick(s socketio.Conn, msg string) {
	log.Println("notice:", msg)
	s.Emit("reply", "have "+msg)
}

func (ws *WebSocket) handleChat(s socketio.Conn, msg string) string {
	s.SetContext(msg)
	return "recv " + msg
}

func (ws *WebSocket) handleBye(s socketio.Conn) string {
	last := s.Context().(string)
	s.Emit("bye", last)
	if err := s.Close(); err != nil {
		return err.Error()
	}
	return last
}

func (ws *WebSocket) handleAction(s socketio.Conn, msg string) {
	ws.Notify(msg)
}

func (ws *WebSocket) handleError(s socketio.Conn, e error) {
	log.Println("meet error:", e)
}

func (ws *WebSocket) handleDisconnect(s socketio.Conn, reason string) {
	log.Println("closed", reason)
}

func (ws *WebSocket) Init() {
	ws.server.OnConnect("/", ws.handleConnect)

	ws.server.OnEvent("/", "notice", ws.handleNotice)

	ws.server.OnEvent("/", "click", ws.handleClick)

	ws.server.OnEvent("/", "action", ws.handleAction)

	ws.server.OnEvent("/chat", "msg", ws.handleChat)

	ws.server.OnEvent("/", "bye", ws.handleBye)

	ws.server.OnError("/", ws.handleError)

	ws.server.OnDisconnect("/", ws.handleDisconnect)

	go func() {
		if err := ws.server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
}

func (ws *WebSocket) Close() {
	ws.server.Close()
}

func (ws *WebSocket) Start(res http.ResponseWriter, req *http.Request) {
	ws.server.ServeHTTP(res, req)
}
