package game

import (
	"strings"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

type rowRenderer struct {
	row      int
	selected []CardIdx
}

func (r rowRenderer) getNthLine(cards []*Card, n int) string {
	cardStrings := make([]string, len(cards))
	for i, card := range cards {
		cardStrings[i] = r.getNthLineForCard(card, n, r.isSelected(i))
	}

	spacing := "  "
	return strings.Join(cardStrings, spacing) + "\n"
}

const cardHeight = 6

func (r rowRenderer) getNthLineForCard(card *Card, n int, isSelected bool) string {
	cardDesign := r.getCardLine(card)

	cardColor := func(s string, i ...interface{}) string {
		return s
	}
	if isSelected {
		cardColor = color.RedString
	}

	switch n {
	case 0:
		return cardColor(" _____ ")
	case 1:
		return cardColor("|     |")
	case 2, 4:
		design := "     "
		if card.number == 1 || card.number == 2 {
			design = cardDesign
		}
		return cardColor("|") + design + cardColor("|")
	case 3:
		design := "     "
		if card.number == 0 || card.number == 2 {
			design = cardDesign
		}
		return cardColor("|") + design + cardColor("|")
	case 5:
		return cardColor("|_____|")
	}
	logrus.Fatalf("Cannot print row %d of card", n)
	return ""
}

func (r rowRenderer) getCardLine(card *Card) string {
	if card.shading < 0 || card.shading > 2 {
		logrus.Fatalf("Invalid shading value of %d", card.shading)
	}
	if card.shape < 0 || card.shape > 2 {
		logrus.Fatalf("Invalid shape value of %d", card.shape)
	}
	shapes := [][]string{
		{"[ ]", "[-]", "[#]"},
		{"( )", "(-)", "(#)"},
		{"< >", "<->", "<#>"},
	}

	var colorFunction func(string, ...interface{}) string
	switch card.color {
	case 0:
		colorFunction = color.HiGreenString
	case 1:
		colorFunction = color.HiBlueString
	case 2:
		colorFunction = color.HiRedString
	default:
		logrus.Fatalf("Invalid shape value of %d", card.shape)
	}

	return colorFunction(" " + shapes[card.shape][card.shading] + " ")
}

func (r rowRenderer) isSelected(col int) bool {
	for _, selectedIdx := range r.selected {
		if selectedIdx.Row == r.row && selectedIdx.Column == col {
			return true
		}
	}
	return false
}
