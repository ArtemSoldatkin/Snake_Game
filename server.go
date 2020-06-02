// websockets.go
package main

import (
	"net/http"
	"snake/field"

	"github.com/gorilla/websocket"
)

// GameField - struct with game field params
var GameField = field.Field{300, 300}

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
			msg = reducer(msg)
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))))
	//http.Handle("/static", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})*/

	http.ListenAndServe(":5000", nil)
}
