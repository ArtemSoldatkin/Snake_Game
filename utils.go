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
