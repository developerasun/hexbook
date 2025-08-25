package auth

import (
	"errors"
	"log"
	"sync"

	// "github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ClientContext struct {
	Connection *websocket.Conn
	SocketID   *uint64
}

type SocketManager struct {
	List      map[uint64]ClientContext
	Count     uint64
	MaxClient uint64
}

func HandleConnection(sc *SocketManager, conn *websocket.Conn) error {
	if sc.Count >= sc.MaxClient {
		error := errors.New("HandleConnection: max socket client reached, please try later.")
		return error
	}

	var rwMutex sync.RWMutex
	rwMutex.Lock()
	sc.List[sc.Count] = ClientContext{Connection: conn, SocketID: &sc.Count}
	sc.Count++
	rwMutex.Unlock()

	log.Println("mutex unlocked and socket connected")
	return nil
}
