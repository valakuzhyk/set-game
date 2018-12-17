package game

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type renderer struct {
	cards    [][]*Card
	selected []CardIdx
}

// CardPrinter returns a string representing cards
type CardPrinter interface {
	PrintCardString()
}

// CreateCardPrinter creates a renderer for the command line
func CreateCardPrinter(cards [][]*Card, selected []CardIdx) CardPrinter {
	if len(cards) != 3 {
		logrus.Fatal("Cards should always be a multiple of three")
	}
	return renderer{cards, selected}
}

func (r renderer) PrintCardString() {
	fmt.Printf(r.getCardString())
}

func (r renderer) getCardString() string {
	verticalSpacing := "\n"

	output := ""
	output += r.getRow(0) + verticalSpacing
	output += r.getRow(1) + verticalSpacing
	output += r.getRow(2)
	return output
}

func (r renderer) getRow(i int) string {
	output := ""

	rowR := rowRenderer{row: i, selected: r.selected}
	for n := 0; n < cardHeight; n++ {
		output += rowR.getNthLine(r.cards[i], n)
	}

	return output
}
