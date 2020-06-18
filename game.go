package main

import (
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

type block struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Game - snake game
type Game struct {
	FieldSize, BlockSize, Speed int
	Direction                   string
	IsStarted                   bool
	done                        chan bool
	ticker                      *time.Ticker
	Snake                       []block
	conn                        *websocket.Conn
}

// Init - initialize game
func (g *Game) Init() {
	g.Direction = "RIGHT"
	g.IsStarted = false
	g.ticker = time.NewTicker(500 * time.Millisecond)
	g.done = make(chan bool)
	g.initSnake()
}

// SetConn - set websocket connection
func (g *Game) SetConn(conn *websocket.Conn) {
	g.conn = conn
}

// SetDirection - set movement direction
func (g *Game) SetDirection(direction string) {
	g.Direction = direction
}

func (g *Game) initSnake() {
	randInt := rand.Intn(int(g.FieldSize/2)) * g.BlockSize
	head := block{randInt, randInt}
	g.Snake = []block{head}
}

func (g *Game) move() {
	head := g.Snake[0]
	switch g.Direction {
	case "UP":
		head.Y -= g.BlockSize
	case "DOWN":
		head.Y += g.BlockSize
	case "LEFT":
		head.X -= g.BlockSize
	case "RIGHT":
		head.X += g.BlockSize
	}
	g.Snake = g.Snake[:len(g.Snake)-1]
	g.Snake = append(g.Snake, head)
}

func (g Game) checkBoundaries() bool {
	head := g.Snake[0]
	max := g.FieldSize * g.BlockSize
	return head.X < 0 || head.X > max-g.BlockSize || head.Y > max-g.BlockSize || head.Y < 0
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
				g.move()
				isOver := g.checkBoundaries()
				var msg []byte
				if isOver {
					msg, _ = createMsg("GAME_OVER", nil)
					g.IsStarted = false
				} else {
					msg, _ = createMsg("MOVE", g.Snake)
				}
				if err := g.conn.WriteMessage(1, msg); err != nil {
					return
				}
				if isOver {
					g.stop()
				}
			}
		}
	}()
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
		g.start()
	} else {
		g.stop()
	}
}
