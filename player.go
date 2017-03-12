package main

import "github.com/nsf/termbox-go"

type player struct {
	x     int
	color termbox.Attribute
	char  rune
}

func (p *player) moveLeft() {
	p.x--
	if p.x < 1 {
		p.x = 1
	}
}

func (p *player) moveRight() {
	p.x++
	if p.x > arenaWidth-1 {
		p.x = arenaWidth - 1
	}
}

func (p *player) fire() {
	bullets[&bullet{char: 'o', x: p.x, y: playerY - 1}] = struct{}{}
}

func (p *player) render() {
	termbox.SetCell(p.x, playerY, p.char, termbox.ColorDefault, termbox.ColorDefault)
}
