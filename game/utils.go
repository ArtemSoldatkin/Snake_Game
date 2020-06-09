package game

import "snake/field"

// StepSize - distance of one step
const StepSize int = 20

// Step - one snake step
func Step(direction string, head field.Block) field.Block {
	x := head.X
	y := head.Y
	h, w := StepSize, StepSize
	switch direction {
	case "UP":
		y -= StepSize
	case "DOWN":
		y += StepSize
	case "RIGHT":
		x += StepSize
	case "LEFT":
		x -= StepSize
	}
	return field.Block{x, y, w, h}
}
