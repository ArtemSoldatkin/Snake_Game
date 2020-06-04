package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
)

// TGameSettings - game settings
type TGameSettings = map[string]int

func getGameSettings(data interface{}) TGameSettings {
	var result = TGameSettings{}
	for k, v := range data.(map[string]interface{}) {
		result[k], _ = strconv.Atoi(v.(string))
	}
	return result
}

// TConfig - type with game params
type TConfig struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Speed  int `json:"speed"`
}

func readConfig() TConfig {
	var config TConfig
	data, _ := ioutil.ReadFile("./static/config.json")
	err := json.Unmarshal(data, &config)
	if err != nil {
		log.Panic(err)
	}
	return config
}
