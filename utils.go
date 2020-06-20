package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Settings - settings from message
type Settings = map[string]int

func readSettingsFromMsg(data interface{}) Settings {
	var result = make(Settings)
	for k, v := range data.(map[string]interface{}) {
		f, _ := v.(float64)
		result[k] = int(f)
	}
	return result
}

// isOpposite - check current direction is opposite to next direction or not
func isOpposite(current, next string) bool {
	switch current {
	case "UP":
		return next == "DOWN"
	case "DOWN":
		return next == "UP"
	case "RIGHT":
		return next == "LEFT"
	case "LEFT":
		return next == "RIGHT"
	default:
		return true
	}
}

//Config - game config from json
type Config struct {
	FieldSize int `json:"field_size"`
	BlockSize int `json:"block_side"`
	Speed     int `json:"speed"`
	MaxSpeed  int `json:"max_speed"`
}

func readConfig() Config {
	var config Config
	data, _ := ioutil.ReadFile("./static/config.json")
	err := json.Unmarshal(data, &config)
	if err != nil {
		log.Panic(err)
	}
	return config
}
