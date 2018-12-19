package main

import (
	"fmt"

	term "github.com/nsf/termbox-go"
	"github.com/valakuzhyk/set-game/game"
)

func main() {
	state := game.Start()
	defer term.Close()
	state.RenderCards()

	fmt.Println("Please enter your choice:")
	for {
		ev := term.PollEvent()
		if ev.Key == term.KeyEsc {
			break
		}
		state.HandleKeyPress(ev)
	}

}
