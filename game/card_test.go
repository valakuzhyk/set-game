package game

import (
	"reflect"
	"testing"
)

func TestCardsAreUnique(t *testing.T) {
	cards := GenDeck()
	for i := range cards {
		for j := i + 1; j < len(cards); j++ {
			if cards[i] == cards[j] {
				t.Fatalf("Cards at %d and %d are the same", i, j)
			}
		}
	}
}

func TestGetSet(t *testing.T) {
	tests := []struct {
		name  string
		cards []*Card
		want  []int
	}{
		{"No cards", []*Card{}, []int{}},
		{"One card", []*Card{{0, 0, 0, 0}}, []int{}},
		{"One difference", []*Card{
			{0, 0, 0, 1},
			{0, 0, 0, 2},
			{0, 0, 0, 3},
		}, []int{0, 1, 2}},
		{"Extra meaningless card", []*Card{
			{0, 0, 0, 1},
			{0, 0, 1, 1},
			{0, 0, 0, 2},
			{0, 0, 0, 3},
		}, []int{0, 2, 3}},
		{"No set", []*Card{
			{0, 0, 0, 1},
			{0, 0, 1, 1},
			{0, 1, 0, 1},
			{1, 0, 0, 1},
		}, []int{}},
		{"Many meaningless cards", []*Card{
			{0, 0, 1, 1},
			{0, 0, 0, 1}, // card1
			{0, 0, 0, 2}, // card2
			{0, 1, 2, 2},
			{2, 2, 2, 2},
			{1, 1, 1, 1},
			{0, 0, 0, 0}, // card3
		}, []int{1, 2, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSet(tt.cards); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasSet(t *testing.T) {
	tests := []struct {
		name  string
		cards []*Card
		want  bool
	}{
		{"No cards", []*Card{}, false},
		{"One card", []*Card{{0, 0, 0, 0}}, false},
		{"One difference", []*Card{
			{0, 0, 0, 1},
			{0, 0, 0, 2},
			{0, 0, 0, 3},
		}, true},
		{"Extra meaningless card", []*Card{
			{0, 0, 0, 1},
			{0, 0, 1, 1},
			{0, 0, 0, 2},
			{0, 0, 0, 3},
		}, true},
		{"No set", []*Card{
			{0, 0, 0, 1},
			{0, 0, 1, 1},
			{0, 1, 0, 1},
			{1, 0, 0, 1},
		}, false},
		{"Many meaningless cards", []*Card{
			{0, 0, 1, 1},
			{0, 0, 0, 1}, // card1
			{0, 0, 0, 2}, // card2
			{0, 0, 2, 2},
			{2, 2, 2, 2},
			{1, 1, 1, 1},
			{0, 0, 0, 3}, // card3
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasSet(tt.cards); got != tt.want {
				t.Errorf("HasSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSet(t *testing.T) {
	type args struct {
		c1 *Card
		c2 *Card
		c3 *Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"One property Different", args{
			&Card{0, 0, 0, 0},
			&Card{0, 0, 0, 1},
			&Card{0, 0, 0, 2},
		}, true},
		{"Not a set", args{
			&Card{0, 0, 0, 0},
			&Card{0, 0, 1, 1},
			&Card{0, 0, 0, 2},
		}, false},
		{"All properties different", args{
			&Card{0, 0, 0, 0},
			&Card{1, 1, 1, 1},
			&Card{2, 2, 2, 2},
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSet(tt.args.c1, tt.args.c2, tt.args.c3); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
