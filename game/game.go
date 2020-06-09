package game

import (
	"fmt"
	"snake/field"
	"time"
)

// Game - type of game with game settings
type Game struct {
	Field     *field.Field
	Speed     int
	IsStarted bool
	Message   func(data interface{})
	ticker    *time.Ticker
	done      chan bool
	Direction string
	BlockSize int
	Snake     []field.Block
}

// Init - initialize game
func (g *Game) Init() {
	g.IsStarted = false
	g.ticker = time.NewTicker(500 * time.Millisecond)
	g.done = make(chan bool)
	g.BlockSize = 20
	head := field.Block{Width: g.BlockSize, Height: g.BlockSize}
	//head.InitRand(g.Field.FieldWidth, g.Field.FieldHeight)
	head.InitRand(350, 350)
	g.Snake = append(g.Snake, head)
	g.Direction = "RIGHT"
}

func (g *Game) start() {
	g.ticker = time.NewTicker(500 * time.Millisecond)
	g.done = make(chan bool)
	go func() {
		for {
			select {
			case <-g.done:
				return
			case <-g.ticker.C:
				newBlock := Step(g.Direction, g.Snake[0])
				g.Snake = g.Snake[:len(g.Snake)-1]
				g.Snake = append(g.Snake, newBlock)
				g.Message(g.Snake)
			}
		}
	}()
}

func (g *Game) stop() {
	g.ticker.Stop()
	g.done <- true
	close(g.done)
	fmt.Println("STOP")
}

// StartStop - start / stop game
func (g *Game) StartStop(isStarted bool) {
	g.IsStarted = isStarted
	if isStarted {
		g.start()
	} else {
		g.stop()
	}
}

// SetSpeed - set game speed
func (g *Game) SetSpeed(speed int) {
	g.Speed = speed
}
