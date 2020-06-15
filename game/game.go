package game

import (
	"snake/field"
	"time"
)

// Game - type of game with game settings
type Game struct {
	Field     *field.Field
	Speed     int
	IsStarted bool
	Message   func(data interface{}, args ...string)
	ticker    *time.Ticker
	done      chan bool
	Direction string
	BlockSize int
	Snake     []field.Block
	gameOver  bool
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
				if checkBorders(g.Field, g.Snake[0]) {
					g.Message(g.Snake)
				} else {
					g.gameOver = true
					g.IsStarted = false
					g.Message(nil, "GAME_OVER")
					g.stop()
				}
			}
		}
	}()
}

func (g *Game) refresh() {
	head := field.Block{Width: g.BlockSize, Height: g.BlockSize}
	head.InitRand(int(g.Field.FieldWidth/2), int(g.Field.FieldHeight/2))
	g.Snake = []field.Block{head}
	g.Direction = "RIGHT"
	g.gameOver = false
}

func (g *Game) stop() {
	g.ticker.Stop()
	g.done <- true
	close(g.done)
}

// StartStop - start / stop game
func (g *Game) StartStop(isStarted bool) {
	g.IsStarted = isStarted
	if isStarted {
		if g.gameOver {
			g.refresh()
		}
		g.start()
	} else {
		g.stop()
	}
}

// SetSpeed - set game speed
func (g *Game) SetSpeed(speed int) {
	g.Speed = speed
}
