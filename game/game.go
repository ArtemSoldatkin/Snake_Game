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
	Message   func()
	ticker    *time.Ticker
	done      chan bool
}

// Init - initialize game
func (g *Game) Init() {
	g.IsStarted = false
	g.ticker = time.NewTicker(500 * time.Millisecond)
	g.done = make(chan bool)
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
				g.Message()
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
