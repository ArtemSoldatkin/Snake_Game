package main

import (
	"encoding/json"
)

// Msg - message type
type Msg struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

func readMsg(data []byte) Msg {
	var msg Msg
	json.Unmarshal(data, &msg)
	return msg
}

func createMsg(action string, data interface{}) ([]byte, error) {
	msg := Msg{Action: action, Data: data}
	return json.Marshal(msg)
}
