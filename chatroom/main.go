package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var mutex sync.RWMutex
var msgq = make(chan []byte, 1000)
var conns = make(map[*websocket.Conn]bool)

func triggerConn(conn *websocket.Conn, status bool) {
	mutex.Lock()
	conns[conn] = status
	mutex.Unlock()
}

func pushMsg() {
	for {
		select {
		case msg := <-msgq:
			fmt.Println(string(msg))
			mutex.RLock()
			for conn, ok := range conns {
				if ok {
					if err := conn.WriteMessage(1, msg); err != nil {
						triggerConn(conn, false)
					}
				}
			}
			mutex.RUnlock()
		}
	}
}

func handleWebsocket(conn *websocket.Conn) {
	triggerConn(conn, true)

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			triggerConn(conn, false)
			fmt.Println("connection closed")
			return
		}
		fmt.Printf("message type: %d, message: %s", messageType, string(msg))
		msgq <- msg
	}
}

type defaultHandler struct{}

func (h *defaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("error in building websocket connection: %v", err)
		return
	}

	go handleWebsocket(conn)
}

func main() {
	go pushMsg()
	http.ListenAndServe("127.0.0.1:8080", &defaultHandler{})
}
