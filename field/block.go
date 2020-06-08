package field

import "math/rand"

// Block - 1 piece of snake
type Block struct {
	X int `json:"x"`
	Y int `json:"y"`
	Width int `json:"width"`
	Height int `json:"height"`
}

// InitRand - set random x & y to block
func (b *Block) InitRand(maxX, maxY int) {
	b.X = rand.Intn(maxX)
	b.Y = rand.Intn(maxY)
}