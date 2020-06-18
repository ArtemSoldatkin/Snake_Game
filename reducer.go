package main

func reducer(message []byte) (result *[]byte) {
	msg := readMsg(message)
	switch msg.Action {

	case "CONNECT":
		game.Init()
		data, _ := createMsg("INITIALIZE", game.Snake)
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
		isStarted := msg.Data.(bool)
		game.StartStop(isStarted)
		data, _ := createMsg("START_PAUSE", nil)
		result = &data

	// Game controls (check currect direction - can't set opposite direction)
	case "UP", "DOWN", "RIGHT", "LEFT":
		game.SetDirection(msg.Action)
		result = nil
	}
	return
}
