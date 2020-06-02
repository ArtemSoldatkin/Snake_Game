// websockets.go
package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			data := readMessage(msg)
			if data.Type == "CONNECT" {
				msg, _ = createMessage("SET_FIELD_SIZE", data.Data)
			}

			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "client.html")
	})

	http.ListenAndServe(":5000", nil)
}
