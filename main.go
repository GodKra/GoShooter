package main

import (
	"github.com/nsf/termbox-go"
	"fmt"
)

func main() {
	e := termbox.Init()
	if e != nil {
		fmt.Print(e)
		return
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputMouse | termbox.InputEsc)
	start()
}
