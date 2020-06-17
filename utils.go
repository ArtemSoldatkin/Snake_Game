package main

import "strconv"

// Settings - settings from message
type Settings = map[string]int

func readSettingsFromMsg(data interface{}) Settings {
	var result Settings
	for k, v := range data.(map[string]interface{}) {
		result[k], _ = strconv.Atoi(v.(string))
	}
	return result
}
