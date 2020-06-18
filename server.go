package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// read config ...
var game = Game{FieldSize: 20, BlockSize: 20, Speed: 60}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		game.SetConn(conn)
		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			message := reducer(msg)
			if message == nil {
				continue
			}
			msg = *message
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe(":5000", nil)
}
