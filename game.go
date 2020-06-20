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

type snakeFood struct {
	Snake []block `json:"snake"`
	Food  []block `json:"food"`
}

// Game - snake game
type Game struct {
	FieldSize, BlockSize, Speed, MaxSpeed int
	Direction                             string
	IsStarted                             bool
	done                                  chan bool
	ticker                                *time.Ticker
	Snake                                 []block
	conn                                  *websocket.Conn
	Food                                  []block
}

// Init - initialize game
func (g *Game) Init() {
	g.Direction = "RIGHT"
	g.IsStarted = false
	//g.ticker = time.NewTicker(200 * time.Millisecond)
	//g.done = make(chan bool)
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

// SetGameParams - set base game parameters
func (g *Game) SetGameParams(fieldSize, blockSize, speed int) {
	g.FieldSize = fieldSize
	g.BlockSize = blockSize
	g.Speed = speed
}

func (g *Game) initFood() {
	randInt := rand.Intn(g.FieldSize-2)*g.BlockSize + g.BlockSize
	food := block{randInt, randInt}
	g.Food = append(g.Food, food)
}

func (g *Game) initSnake() {
	randInt := rand.Intn(int(g.FieldSize/2))*g.BlockSize + g.BlockSize
	head := block{randInt, randInt}
	g.Snake = []block{head}
}

func (g *Game) isEatFood(head block) bool {
	for i, f := range g.Food {
		if f.X == head.X && f.Y == head.Y {
			if len(g.Food) > 1 {
				g.Food[i] = g.Food[len(g.Food)-1]
				g.Food = g.Food[:len(g.Food)-1]
			} else {
				g.Food = []block{}
			}
			return true
		}
	}
	return false
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
	if !g.isEatFood(head) {
		g.Snake = g.Snake[:len(g.Snake)-1]
	}
	g.Snake = append([]block{head}, g.Snake...)
}

func (g Game) isSelfBite() bool {
	if len(g.Snake) == 1 {
		return false
	}
	head := g.Snake[0]
	for _, p := range g.Snake[1:] {
		if p.X == head.X && p.Y == head.Y {
			return true
		}
	}
	return false
}

func (g Game) checkBoundaries() bool {
	head := g.Snake[0]
	max := g.FieldSize * g.BlockSize
	return head.X < 0 || head.X > max-g.BlockSize || head.Y > max-g.BlockSize || head.Y < 0 || g.isSelfBite()
}

func (g *Game) start() {
	g.ticker = time.NewTicker(time.Duration(((g.MaxSpeed + 1 - g.Speed) * 100)) * time.Millisecond)
	g.done = make(chan bool)
	go func() {
		for {
			select {
			case <-g.done:
				return
			case <-g.ticker.C:
				if len(g.Food) == 0 {
					g.initFood()
				}
				g.move()
				isOver := g.checkBoundaries()
				var msg []byte
				if isOver {
					msg, _ = createMsg("GAME_OVER", nil)
					g.IsStarted = false
				} else {
					msg, _ = createMsg("MOVE", snakeFood{g.Snake, g.Food})
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
