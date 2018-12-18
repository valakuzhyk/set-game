package game

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

// CardIdx is used to represent a card on the field.
type CardIdx struct {
	Row    int
	Column int
}

// State represents the current state of the game
type State struct {
	field [][]*Card
	deck  []*Card

	selected []CardIdx

	score int
}

func (s State) String() string {
	output := fmt.Sprintln("Field: ")
	for _, row := range s.field {
		for _, card := range row {
			output += fmt.Sprintf("%+v", card)
		}
		output += "\n"
	}
	output += "\n"

	output += fmt.Sprintf("Num cards in deck: %d\n", len(s.deck))
	output += fmt.Sprintf("Selected: %v\n", s.selected)
	output += fmt.Sprintf("Score: %d\n", s.score)
	return output
}

// Start creates a new game state to play with.
func Start() State {
	err := term.Init()
	if err != nil {
		panic(err)
	}

	deck := GenDeck()
	Shuffle(deck)

	fmt.Println(`
Welcome! 

To select cards, use your keyboard

q w e r ...
a s d f ...
z x c v ...

Try and get the most cards that make sets!
If you can't find a set, press spacebar to get more cards on the field.

Press ESC to exit

`)
	term.PollEvent()

	// Start out with 12 cards on the field
	return State{
		field: [][]*Card{
			append([]*Card{}, deck[:4]...),
			append([]*Card{}, deck[4:8]...),
			append([]*Card{}, deck[8:12]...),
		},
		deck: deck[12:],
	}
}

// WaitForKey returns an event representing the key press
func (s State) WaitForKey() term.Event {
	return term.PollEvent()
}

// RenderCards prints the cards out on the command line.
func (s State) RenderCards() {
	term.Sync()
	cardPrinter := CreateCardPrinter(s.field, s.selected)
	cardPrinter.PrintCardString()
	fmt.Printf("Your score: %d\n", s.score)
}

// ClearSelections removes all selections.
func (s *State) ClearSelections() {
	s.selected = []CardIdx{}
}

// AddColumn adds a card to each row, creating a new column.
func (s *State) AddColumn() {
	for i := 0; i < 3; i++ {
		if len(s.deck) == 0 {
			break
		}

		s.field[i] = append(s.field[i], s.deck[0])
		s.deck = s.deck[1:]
	}
	s.RenderCards()
	return
}

// Select chooses a card
func (s *State) Select(idx CardIdx) {
	for i, val := range s.selected {
		if val == idx {
			s.selected = append(s.selected[:i], s.selected[i+1:]...)
			s.RenderCards()
			return
		}
	}
	s.selected = append(s.selected, idx)
	if len(s.selected) != 3 {
		s.RenderCards()
		return
	}
	if s.CheckSet(s.selected[0], s.selected[1], s.selected[2]) {
		s.score++
		s.drawNewCard(s.selected...)
		s.selected = []CardIdx{}
		s.RenderCards()
		fmt.Println("Congrats! You got a set")
	} else {
		s.RenderCards()
		fmt.Println("That's not a set :|")
		s.selected = []CardIdx{}
	}
}

// CheckSet returns whether or not the cards at the specified indices make up a set.
// If so, they are removed, and enough cards are added to put the field back to 12 again (if possible)
func (s State) CheckSet(i1, i2, i3 CardIdx) bool {
	if !IsSet(s.getCard(i1), s.getCard(i2), s.getCard(i3)) {
		return false
	}
	return true
}

// HasSet returns true if the field has a set
func (s State) HasSet() bool {
	return len(s.GetSet()) > 0
}

// GetSet returns a slice of indices representing a set of cards if a set exists in the field.
func (s State) GetSet() []CardIdx {
	cards := []*Card{}
	for _, row := range s.field {
		cards = append(cards, row...)
	}
	indices := GetSet(cards)

	cardIndices := []CardIdx{}

outerLoop:
	for _, idx := range indices {
		currentIdx := 0
		for row, cardsInRow := range s.field {
			for col := range cardsInRow {
				if currentIdx == idx {
					cardIndices = append(cardIndices, CardIdx{Row: row, Column: col})
					continue outerLoop
				}
				currentIdx++
			}
		}
	}
	return cardIndices
}

// draws new cards at the selected indices, if possible.
func (s *State) drawNewCard(indices ...CardIdx) {
	for _, idx := range indices {
		row, col := idx.Row, idx.Column
		if len(s.deck) == 0 {
			s.field[row] = append(s.field[row][:col], s.field[row][col+1:]...)
			continue
		}
		newCard := s.deck[0]
		s.deck = s.deck[1:]
		s.field[row][col] = newCard
	}
}

func (s State) getCard(idx CardIdx) *Card {
	return s.field[idx.Row][idx.Column]
}
