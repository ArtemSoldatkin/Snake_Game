package main

func reducer(msg []byte) (result []byte) {
	message := readMessage(msg)
	switch message.Type {
	case "CONNECT":
		size := [2]int{GameField.Width, GameField.Height}
		result, _ = createMessage("SET_FIELD_SIZE", size)
	}
	return
}
