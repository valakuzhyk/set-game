package main

import (
	"fmt"

	term "github.com/nsf/termbox-go"
	"github.com/valakuzhyk/set-game/game"
)

var enableGodMode = false

// func reset() {
// 	term.Sync() // cosmestic purpose
// }

var keyMap = map[string]game.CardIdx{
	"q": game.CardIdx{Row: 0, Column: 0}, "a": game.CardIdx{Row: 1, Column: 0}, "z": game.CardIdx{Row: 2, Column: 0},
	"w": game.CardIdx{Row: 0, Column: 1}, "s": game.CardIdx{Row: 1, Column: 1}, "x": game.CardIdx{Row: 2, Column: 1},
	"e": game.CardIdx{Row: 0, Column: 2}, "d": game.CardIdx{Row: 1, Column: 2}, "c": game.CardIdx{Row: 2, Column: 2},
	"r": game.CardIdx{Row: 0, Column: 3}, "f": game.CardIdx{Row: 1, Column: 3}, "v": game.CardIdx{Row: 2, Column: 3},
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
			switch string(ev.Ch) {
			case "1":
				fmt.Println(state)
			case "2":
				if state.HasSet() {
					fmt.Println("There is a set present")
				} else {
					fmt.Println("There is no set present :O")
				}
			case "3":
				fmt.Println(state.GetSet())
			case "4":
				state.ClearSelections()
				for _, selection := range state.GetSet() {
					state.Select(selection)
				}
				fmt.Println("You lazy bum...")
			}
		}

	}
	term.Close()
}
