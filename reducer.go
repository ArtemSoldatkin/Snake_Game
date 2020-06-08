package main

import (
	"fmt"
)

func reducer(msg []byte) (result []byte) {
	message := readMessage(msg)
	switch message.Type {
	case "CONNECT":
		GameField.Init()
		head := *GameField.Head
		result, _ = createMessage("INITIALIZE", head)

	// Game config
	case "SET_GAME_SETTINGS":
		settings := getGameSettings(message.Data)
		GameField.SetSize(settings["width"], settings["height"])
		result, _ = createMessage("SET_GAME_SETTINGS", nil)
	case "START_PAUSE":
		isStarted := message.Data.(bool)
		Game.StartStop(isStarted)
		result, _ = createMessage("START_PAUSE", isStarted)

	// Game controls
	case "UP":
		fmt.Println("UP")
		result, _ = createMessage("UP", "UP")

	case "DOWN":
		fmt.Print("DOWN")
		result, _ = createMessage("DOWN", "DOWN")
	case "RIGHT":
		fmt.Print("RIGHT")
		result, _ = createMessage("RIGHT", "RIGHT")

	case "LEFT":
		fmt.Print("LEFT")
		result, _ = createMessage("LEFT", "LEFT")

	}
	return
}
