package messenger

import (
	"github.com/devplayg/hippo"
	"github.com/devplayg/messenger/entity"
)

func NewHub() *Hub {
	return &Hub{
		rooms:      nil,
		Launcher:   hippo.Launcher{},
		register:   make(chan *entity.ChatRoom),
		unregister: make(chan *entity.ChatRoom),
	}

}

type Hub struct {
	rooms map[*entity.ChatRoom]bool
	hippo.Launcher
	register   chan *entity.ChatRoom
	unregister chan *entity.ChatRoom
}

func (h *Hub) Start() error {
	for {
		select {
		case room := <-h.register:
			h.rooms[room] = true
		case room := <-h.unregister:
			delete(h.rooms, room)
		case <-h.Ctx.Done():
			return nil
		}
	}
}

func (h *Hub) Stop() error {
	return nil
}

type ChatRoom struct {
	members    map[*entity.Member]bool
	broadcast  chan *entity.Message
	register   chan *entity.Member
	unregister chan *entity.Member
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		broadcast:  make(chan *entity.Message),
		register:   make(chan *entity.Member),
		unregister: make(chan *entity.Member),
		members:    make(map[*entity.Member]bool),
	}
}
