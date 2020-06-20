package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// CONFIG - game config from json
var CONFIG = readConfig()

var game = Game{FieldSize: CONFIG.FieldSize, BlockSize: CONFIG.BlockSize, Speed: CONFIG.Speed, MaxSpeed: CONFIG.MaxSpeed}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
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
