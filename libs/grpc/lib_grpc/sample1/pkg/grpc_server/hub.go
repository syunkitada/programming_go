package grpc_server

import (
	"github.com/golang/glog"

	"sample1/pkg/grpc_pb"
)

type Hub struct {
	// Registered clients.
	streams map[grpc_pb.Simple_ChatServer]bool

	// Inbound messages from the clients.
	broadcast chan *grpc_pb.ChatStream

	// Register requests from the clients.
	register chan grpc_pb.Simple_ChatServer

	// Unregister requests from clients.
	unregister chan grpc_pb.Simple_ChatServer
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan *grpc_pb.ChatStream),
		register:   make(chan grpc_pb.Simple_ChatServer),
		unregister: make(chan grpc_pb.Simple_ChatServer),
		streams:    make(map[grpc_pb.Simple_ChatServer]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			glog.Info("registered: %v", client)
			h.streams[client] = true
		case client := <-h.unregister:
			if _, ok := h.streams[client]; ok {
				delete(h.streams, client)
				glog.Info("unregistered: %v", client)
			}
		case message := <-h.broadcast:
			for client := range h.streams {
				if err := client.Send(message); err != nil {
					glog.Error(err)
					delete(h.streams, client)
				}
			}
		}
	}

}
