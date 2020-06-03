package main

import "strconv"

// TGameSettings - game settings
type TGameSettings = map[string]int

func getGameSettings(data interface{}) TGameSettings {
	var result = TGameSettings{}
	for k, v := range data.(map[string]interface{}) {
		result[k], _ = strconv.Atoi(v.(string))
	}
	return result
}
