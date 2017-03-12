package main

import "github.com/nsf/termbox-go"

type bullet struct {
	char  rune
	x, y  int
}

func (b *bullet) move() {
	b.y--
	if b.y < 0 {
		delete(bullets, b)
	}
}

func (b *bullet) render() {
	termbox.SetCell(b.x, b.y, b.char, termbox.ColorRed, termbox.ColorDefault)
}

func (b *bullet) update() {
	b.move()
}
