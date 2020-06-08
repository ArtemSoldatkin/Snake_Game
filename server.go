// websockets.go
package main

import (
	"net/http"
	"snake/field"
	"snake/game"

	"github.com/gorilla/websocket"
)

// GameConfig - default game config
var GameConfig = readConfig()

// GameField - struct with game field params
var GameField = field.Field{FieldWidth: GameConfig.Width, FieldHeight: GameConfig.Height}

// Game - type of game with settings
var Game = game.Game{Field: &GameField, Speed: GameConfig.Speed}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {

	Game.Init()

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		moveAction := func() {
			msg, _ := createMessage("MOVE", nil)
			if err := conn.WriteMessage(1, msg); err != nil {
				return
			}
		}
		Game.Message = moveAction
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
	http.ListenAndServe(":5000", nil)
}
