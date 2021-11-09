package hub

import (
	"fmt"

	"github.com/midoks/go-p2p-server/internal/client"
	"github.com/midoks/go-p2p-server/internal/hub/cmap"
)

var h *Hub

type Hub struct {
	Clients cmap.ConMap
}

func Init() {
	h = &Hub{
		Clients: cmap.New(),
	}
}

func GetInstance() *Hub {
	return h
}

func DoRegister(client *client.Client) {
	fmt.Println("hub DoRegister:", client.PeerId)
	if client.PeerId != "" {
		h.Clients.Set(client.PeerId, client)
	} else {
		panic("DoRegister")
	}
}
func GetClientNum() int {
	return h.Clients.CountNoLock()
}

func RemoveClient(id string) {
	h.Clients.Remove(id)
}

func DoUnregister(peerId string) bool {
	fmt.Println("hub DoUnregister :", peerId)
	if peerId == "" {
		return false
	}
	if h.Clients.Has(peerId) {
		h.Clients.Remove(peerId)
		return true
	}
	return false
}