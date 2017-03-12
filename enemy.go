package main

import "github.com/nsf/termbox-go"

type enemy struct {
	x, y int
	char rune
}

func (e *enemy) move() {
	e.y++
	if e.y >= arenaHeight {
		gameOver = true
	}
}

func (e *enemy) render() {
	termbox.SetCell(e.x, e.y, e.char, termbox.ColorRed, termbox.ColorDefault)
}

func (e *enemy) update() {
	e.move()
}

func spawnEnemy(x, y int) {
	enemies[&enemy{x, y, 'W'}] = struct{}{}
}