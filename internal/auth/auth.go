package auth

import (
	"errors"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type ClientContext struct {
	Connection *websocket.Conn
	SocketID   int
}

type SocketManager struct {
	List      map[int]*ClientContext
	Count     int
	MaxClient int
	Mutex     sync.RWMutex
}

func HandleConnection(sc *SocketManager, conn *websocket.Conn) error {
	sc.Mutex.Lock()
	defer sc.Mutex.Unlock()

	if sc.Count >= sc.MaxClient {
		error := errors.New("HandleConnection: max socket client reached, please try later")
		return error
	}

	sc.Count++
	sc.List[sc.Count] = &ClientContext{Connection: conn, SocketID: sc.Count}

	log.Println("mutex unlocked and socket connected")
	return nil
}

func GetConnection(sc *SocketManager, id int) (*websocket.Conn, error) {
	sc.Mutex.RLock()
	defer sc.Mutex.RUnlock()

	hasContext := sc.List[id]
	if hasContext == nil {
		error := errors.New("GetConnection: no such client")
		log.Println(error.Error())
		return nil, error
	}
	conn := sc.List[id].Connection

	return conn, nil
}
