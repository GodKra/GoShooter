package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
	"fmt"
)

const (
	playerY     = arenaHeight - 1
	arenaHeight = 20
	arenaWidth  = 50
)

var (
	bullets = make(map[*bullet]struct{}, 0)
	p       = player{arenaHeight - 1, termbox.ColorGreen, 'A'}
	enemies = make(map[*enemy]struct{}, 0)

	enemyTicker = time.NewTicker(time.Millisecond * time.Duration(enemyDelay))
	spawnTick   = time.Tick(time.Second * 5)
	bulletTick  = time.Tick(time.Millisecond * 100)

	lastFired  time.Time
	enemyDelay int64 = 800
	randGen          = rand.New(rand.NewSource(time.Now().UnixNano()))
	gameOver   bool
	score  =  0
)

type renderable interface {
	render()
}

type updatable interface {
	update()
}

type entity interface {
	updatable
	renderable
}

func start() {
	cChan := make(chan command)
	go handleEvents(cChan)
	spawnEnemy(randomCoords())
	drawScore()
	mainLoop(cChan)
	gameOverScreen()
}

func mainLoop(cChan chan command) {
	for {
		select {
		case c := <-cChan:
			switch c {
			case left:
				p.moveLeft()
			case right:
				p.moveRight()
			case fire:
				if time.Now().Sub(lastFired) > time.Millisecond*250 {
					p.fire()
					lastFired = time.Now()
				}

			case exit:
				return
			}
		case <-bulletTick:
			bulletLoop:
			for b := range bullets {
				b.update()
				for e := range enemies {
					if b.x == e.x && b.y <= e.y {
						collided(b, e)
						continue bulletLoop
					}
				}
			}
		case <-enemyTicker.C:
			for e := range enemies {
				e.update()
			}
		case <-spawnTick:
			spawnEnemy(randomCoords())
		default:
			if gameOver {
				return
			}
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			p.render()
			for b := range bullets {
				b.render()
			}
			for e := range enemies {
				e.render()
			}
			drawBorder()
			termbox.Flush()
			time.Sleep(time.Millisecond * 16)
		}
	}
}

func drawBorder() {
	termbox.SetCell(0, 0, '┏', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(arenaWidth, 0, '┓', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(0, arenaHeight, '┗', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(arenaWidth, arenaHeight, '┛', termbox.ColorDefault, termbox.ColorDefault)

	//horizontal
	for x := 1; x < arenaWidth; x++ {
		termbox.SetCell(x, 0, '━', termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x, arenaHeight, '━', termbox.ColorDefault, termbox.ColorDefault)
	}
	//vertical
	for y := 1; y < arenaHeight; y++ {
		termbox.SetCell(0, y, '┃', termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(arenaWidth, y, '┃', termbox.ColorDefault, termbox.ColorDefault)
	}
}

func randomCoords() (x, y int) {
	return randGen.Intn(arenaWidth), 0
}

func drawScore() {
	termbox.SetCursor(0,arenaHeight + 1)
	fmt.Printf("Score: %v", score)
	termbox.HideCursor()
}

func gameOverScreen() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawBorder()
	termbox.SetCursor(arenaWidth/2-10, arenaHeight/2)
	fmt.Print("Game over")
	termbox.SetCursor(arenaWidth/2-10, arenaHeight/2+1)
	fmt.Printf("Score: %v", score)
	termbox.HideCursor()
	time.Sleep(time.Second * 2)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func collided(b *bullet, e *enemy) {
	score++
	delete(enemies, e)
	delete(bullets, b)
	enemyTicker.Stop()
	enemyDelay -= 100
	if enemyDelay <= 100 {
		enemyDelay = 100
	}
	enemyTicker = time.NewTicker(time.Millisecond * time.Duration(enemyDelay))
	drawScore()
}