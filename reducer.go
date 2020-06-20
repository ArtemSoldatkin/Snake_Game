package main

func reducer(message []byte) (result *[]byte) {
	msg := readMsg(message)
	switch msg.Action {

	case "CONNECT":
		game.Init()
		data, _ := createMsg("INITIALIZE", snakeFood{Snake: game.Snake})
		result = &data

	// Game config
	case "SET_GAME_SETTINGS":
		// error
		data := readSettingsFromMsg(msg.Data)
		game.SetGameParams(data["field_size"], data["block_side"], data["speed"])
		result = nil

	case "START_PAUSE":
		isStarted := msg.Data.(bool)
		game.StartStop(isStarted)
		data, _ := createMsg("START_PAUSE", nil)
		result = &data

	// Game controls
	case "UP", "DOWN", "RIGHT", "LEFT":
		if !isOpposite(game.Direction, msg.Action) {
			game.SetDirection(msg.Action)
		}
		result = nil
	}
	return
}
