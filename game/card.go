package game

import (
	"math"
	"time"

	"golang.org/x/exp/rand"
)

// Cards have four properties, each with three values
//  Shape:  circle, diamond, square
//  Number: one, two, three
//  Color:  red, blue green
//  Shading:  empty, partial, full

// Card represents a card in the game.
type Card struct {
	shape   int
	number  int
	color   int
	shading int
}

// GenDeck returns the deck of cards that can be used to play a game of set.
func GenDeck() []*Card {
	deck := make([]*Card, int(math.Pow(3, 4)))
	for i := 0; i < int(math.Pow(3, 4)); i++ {
		deck[i] = &Card{
			(i / 1) % 3,
			(i / 3) % 3,
			(i / 9) % 3,
			(i / 27) % 3,
		}
	}

	return deck
}

// Shuffle shuffles the given cards
func Shuffle(cards []*Card) {
	rand.Seed(uint64(time.Now().UnixNano()))
	// Fisher - Yates Shuffle
	for i := len(cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
}

// HasSet returns true if three cards make a set within the cards given.
func HasSet(cards []*Card) bool {
	for i, ci := range cards {
		for j, cj := range cards[i+1:] {
			for _, ck := range cards[j+1:] {
				if IsSet(ci, cj, ck) {
					return true
				}
			}
		}
	}
	return false
}

// IsSet returns whether 3 cards compromise a set
func IsSet(c1, c2, c3 *Card) bool {
	return PropertyAllSameOrAllDifferent(c1.shape, c2.shape, c3.shape) &&
		PropertyAllSameOrAllDifferent(c1.number, c2.number, c3.number) &&
		PropertyAllSameOrAllDifferent(c1.color, c2.color, c3.color) &&
		PropertyAllSameOrAllDifferent(c1.shading, c2.shading, c3.shading)
}

// PropertyAllSameOrAllDifferent returns true if the three numbers are all the same or all different.
func PropertyAllSameOrAllDifferent(p1, p2, p3 int) bool {
	if p1 == p2 && p2 == p3 {
		return true
	}
	if p1 != p2 && p2 != p3 && p1 != p3 {
		return true
	}
	return false
}
