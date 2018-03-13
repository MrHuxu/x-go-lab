package main

import (
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var mutex sync.RWMutex
var conns = make(map[*websocket.Conn]bool)

func openConn(conn *websocket.Conn) {
	mutex.Lock()
	conns[conn] = true
	mutex.Unlock()
}

func closeConn(conn *websocket.Conn) {
	mutex.Lock()
	delete(conns, conn)
	mutex.Unlock()
}

type message struct {
	addr    net.Addr
	content []byte
}

var msgq = make(chan message, 1000)

func pushMsg() {
	for {
		select {
		case msg := <-msgq:
			log.Printf("message from %v received and broadcasted: %s\n", msg.addr, string(msg.content))
			mutex.RLock()
			for conn := range conns {
				if err := conn.WriteMessage(1, msg.content); err != nil {
					closeConn(conn)
				}
			}
			mutex.RUnlock()
		}
	}
}

func handleWebsocket(conn *websocket.Conn) {
	openConn(conn)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			closeConn(conn)
			log.Println("connection closed")
			return
		}
		msgq <- message{addr: conn.RemoteAddr(), content: msg}
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error in building websocket connection: %v", err)
		return
	}

	go handleWebsocket(conn)
}

func main() {
	go pushMsg()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/chat", serveWS)
	http.ListenAndServe(":8080", nil)
}
