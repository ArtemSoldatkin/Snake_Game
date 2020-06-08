package main

import (
	"encoding/json"
)

// message - message to client
type message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func readMessage(data []byte) message {
	var msg message
	json.Unmarshal(data, &msg)
	return msg
}

func createMessage(msgType string, data interface{}) ([]byte, error) {
	msg := message{Type: msgType, Data: data}
	return json.Marshal(msg)
}
