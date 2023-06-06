package manager

import (
	"im-api/im-ws/common/e"
	"sync"

	"github.com/gorilla/websocket"
)

type connectionManager struct {
	// key:在线用户id, value:websocket连接
	sockets map[uint32]*websocket.Conn
	lock *sync.RWMutex
}

var connManager = &connectionManager {
	sockets: make(map[uint32]*websocket.Conn, 0),
	lock: new(sync.RWMutex),
}

func Put(id uint32, conn *websocket.Conn) error {
	if _, ok := connManager.sockets[id]; ok {
		return e.AlreadyExistConnErr
	}

	connManager.lock.Lock()
	defer connManager.lock.Unlock()
	connManager.sockets[id] = conn
	return nil
}

func Get(id uint32) *websocket.Conn {
	connManager.lock.RLock()
	defer connManager.lock.RUnlock()
	return connManager.sockets[id]
}

func Exist(id uint32) bool {
	connManager.lock.RLock()
	defer connManager.lock.RUnlock()
	_, ok := connManager.sockets[id]
	return ok
}

func Remove(id uint32) {
	connManager.lock.Lock()
    defer connManager.lock.Unlock()
    delete(connManager.sockets, id)
}

func Send(id uint32, data []byte) error {
	connManager.lock.Lock()
	defer connManager.lock.Unlock()
	conn, ok := connManager.sockets[id]
	if ok {
		err := conn.WriteMessage(websocket.TextMessage, data)
		return err
	} else {
		return e.NotExistConnErr
	}
}