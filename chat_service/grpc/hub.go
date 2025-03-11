package main

import (
	"log"
	"sync"

	"github.com/insanXYZ/proto/gen/go/chat"
)

type id = string

type Hub struct {
	Clients   map[id](*Client)
	Broadcast chan *chat.MessageResponse
	Register  chan *Client
	mtx       sync.Mutex
}

func NewHub() *Hub {
	h := &Hub{
		Clients:   make(map[id]*Client),
		Broadcast: make(chan *chat.MessageResponse),
		Register:  make(chan *Client),
	}
	go h.Run()
	return h
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mtx.Lock()
			h.Append(client.user.Id, client)
			h.mtx.Unlock()
		case message := <-h.Broadcast:
			for _, v := range h.Clients {
				err := v.stream.Send(message)
				if err != nil {
					log.Println("Error sending message to ", v.user.Name)
					h.Pop(v.user.Id)
				}
			}
		}
	}
}

func (h *Hub) Append(key id, client *Client) {
	if _, ok := h.Clients[key]; !ok {
		h.Clients[key] = client
	}
}

func (h *Hub) Pop(id id) {
	if _, ok := h.Clients[id]; ok {
		delete(h.Clients, id)
	}
}

func (h *Hub) ExistClient(id id) bool {
	_, ok := h.Clients[id]
	return ok
}
