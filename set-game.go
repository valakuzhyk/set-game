package main

import (
	"fmt"

	term "github.com/nsf/termbox-go"
	"github.com/valakuzhyk/set-game/game"
)

var enableGodMode = true

// func reset() {
// 	term.Sync() // cosmestic purpose
// }

var keyMap = map[string]int{
	"q": 0, "a": 1, "z": 2,
	"w": 3, "s": 4, "x": 5,
	"e": 6, "d": 7, "c": 8,
	"r": 9, "f": 10, "v": 11,
}

func main() {
	state := game.Start()
	state.RenderCards()

	fmt.Println("Please enter your choice:")
	for {
		ev := term.PollEvent()
		if ev.Key == term.KeyEsc {
			break
		}
		if idx, ok := keyMap[string(ev.Ch)]; ok {
			state.Select(idx)
		}

		if enableGodMode {
			if string(ev.Ch) == "1" {
				fmt.Println(state)
			}
			if string(ev.Ch) == "2" {
				if state.HasSet() {
					fmt.Println("There is a set present")
				} else {
					fmt.Println("There is no set present :O")
				}
			}
		}

	}
	term.Close()
}
