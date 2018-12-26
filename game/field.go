package game

import (
	"fmt"
)

// Field comprises the cards that are visible to the players.
type Field [][]*Card

func (f Field) String() string {
	output := ""
	for _, row := range f {
		for _, card := range row {
			output += fmt.Sprintf("%+v", card)
		}
		output += "\n"
	}
	return output
}

// CreateField creates an evenly distributed field with the given cards.
func CreateField(cards []*Card) Field {
	field := Field(make([][]*Card, 3))
	oneThird := (2 + len(cards)) / 3
	twoThirds := oneThird + (1+len(cards))/3
	field[0] = append(field[0], cards[0:oneThird]...)
	field[1] = append(field[1], cards[oneThird:twoThirds]...)
	field[2] = append(field[2], cards[twoThirds:]...)

	return field
}

// RedistributeCards ensures that rows are as even as possible.
func (f *Field) RedistributeCards() {
	cards := []*Card{}
	for i, row := range *f {
		for j := range row {
			card := (*f)[i][j]
			if card == nil {
				continue
			}
			cards = append(cards, card)
		}
	}

	*f = CreateField(cards)
}

// NumColumns returns the maximum number of columns of the rows in the field.
func (f Field) NumColumns() int {
	numColumns := 0
	for _, row := range f {
		if len(row) > numColumns {
			numColumns = len(row)
		}
	}
	return numColumns
}

// ReplaceCardAt replaces with a new card. If the new card is nil, this ReplaceCardAt removes the card
// at the given index
func (f *Field) ReplaceCardAt(idx CardIdx, newCard *Card) {
	row, col := idx.Row, idx.Column
	(*f)[row][col] = newCard
}
