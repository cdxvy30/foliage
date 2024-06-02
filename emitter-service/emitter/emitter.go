package emitter

import (
	SocketIO "github.com/yosuke-furukawa/socket.io-go-emitter"
)

func CreateClient() *SocketIO.Emitter {
	emitter, _ := SocketIO.NewEmitter(&SocketIO.EmitterOpts{
		Host: "localhost",
		Port: 6379,
	})
	return emitter
}
