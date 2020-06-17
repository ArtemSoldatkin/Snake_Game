package main

func reducer(message []byte) (result *[]byte) {
	msg := readMsg(message)
	switch msg.Action {

	case "CONNECT":
		data, _ := createMsg("INITIALIZE", nil)
		result = &data

	// Game config
	case "SET_GAME_SETTINGS":
		// error
		data := readSettingsFromMsg(msg.Data)

		game.FieldSize = data["fieldSize"]
		game.BlockSize = data["blockSize"]
		game.Speed = data["speed"]
		result = nil

	case "START_PAUSE":
		data, _ := createMsg("START_PAUSE", nil)
		result = &data

	// Game controls
	case "UP":
		game.Direction = "UP"
		result = nil

	case "DOWN":
		game.Direction = "DOWN"
		result = nil

	case "RIGHT":
		game.Direction = "RIGHT"
		result = nil

	case "LEFT":
		game.Direction = "LEFT"
		result = nil
	}
	return
}
