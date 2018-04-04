package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebsocket(conn *websocket.Conn) {
	ticker := time.Tick(time.Second)

	for {
		select {
		case <-ticker:
			conn.WriteMessage(1, []byte(time.Now().String()+"\n"))
		}
	}
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
	http.HandleFunc("/ws", serveWS)
	http.ListenAndServe(":8080", nil)
}
