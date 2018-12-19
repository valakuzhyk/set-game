package game

import "fmt"

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
	for i, card := range cards {
		row := i % 3
		field[row] = append(field[row], card)

	}
	return field
}

// RedistributeCards ensures that rows are as even as possible.
func (f *Field) RedistributeCards() {
	cards := []*Card{}
	for _, row := range *f {
		cards = append(cards, row...)
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
	if newCard != nil {
		(*f)[row][col] = newCard
	} else {
		(*f)[row] = append((*f)[row][:col], (*f)[row][col+1:]...)
	}
}
