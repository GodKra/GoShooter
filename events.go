package main

import "github.com/nsf/termbox-go"

type command byte

const (
	left command = iota
	right
	fire
	exit
)

func handleEvents(c chan command) {
	for {
		switch e := termbox.PollEvent(); e.Type {
		case termbox.EventKey:
			switch e.Key {
			case termbox.KeyArrowLeft:
				c <- left
			case termbox.KeyArrowRight:
				c <- right
			case termbox.KeySpace:
				c <- fire
			case termbox.KeyEsc:
				c <- exit
			}
		}
	}
}
