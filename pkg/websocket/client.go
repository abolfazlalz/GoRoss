package websocket

import "github.com/gorilla/websocket"

type Client struct {
	ID uint
	ws *websocket.Conn
}

type Clients struct {
	clients []*Client
	lastId  uint
}

func (c Clients) Add(client *Client) {
	c.lastId++
	client.ID = c.lastId
	c.clients = append(c.clients, client)
}
