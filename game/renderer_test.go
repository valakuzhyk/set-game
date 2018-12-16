package game

import (
	"testing"
)

func Test_getRow(t *testing.T) {
	tests := []struct {
		name  string
		cards []*Card
		want  string
	}{
		{"1 Card", []*Card{{0, 0, 0, 0}},
			"" +
				" _____ \n" +
				"|     |\n" +
				"|     |\n" +
				"| [ ] |\n" +
				"|     |\n" +
				"|_____|\n",
		},
		{"2 Cards", []*Card{{0, 0, 0, 0}, {0, 0, 0, 1}},
			"" +
				" _____    _____ \n" +
				"|     |  |     |\n" +
				"|     |  |     |\n" +
				"| [ ] |  | [-] |\n" +
				"|     |  |     |\n" +
				"|_____|  |_____|\n",
		},
		{"3 Cards", []*Card{{0, 0, 0, 0}, {0, 0, 0, 1}, {0, 1, 0, 0}},
			"" +
				" _____    _____    _____ \n" +
				"|     |  |     |  |     |\n" +
				"|     |  |     |  | [ ] |\n" +
				"| [ ] |  | [-] |  |     |\n" +
				"|     |  |     |  | [ ] |\n" +
				"|_____|  |_____|  |_____|\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := renderer{cards: [][]*Card{
				tt.cards,
				{},
				{}}}
			if got := r.getRow(0); got != tt.want {
				t.Errorf("getRow() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

func Test_getCardOutputString(t *testing.T) {
	tests := []struct {
		name  string
		cards [][]*Card
		want  string
	}{
		{"one column", [][]*Card{
			{{0, 0, 0, 0}},
			{{0, 0, 0, 1}},
			{{0, 0, 0, 2}}},
			"" +
				" _____ \n" +
				"|     |\n" +
				"|     |\n" +
				"| [ ] |\n" +
				"|     |\n" +
				"|_____|\n" +
				"\n" +
				" _____ \n" +
				"|     |\n" +
				"|     |\n" +
				"| [-] |\n" +
				"|     |\n" +
				"|_____|\n" +
				"\n" +
				" _____ \n" +
				"|     |\n" +
				"|     |\n" +
				"| [#] |\n" +
				"|     |\n" +
				"|_____|\n",
		},
		{"two column", [][]*Card{
			{{0, 0, 0, 0}, {0, 0, 0, 1}},
			{{0, 0, 0, 2}, {0, 1, 0, 0}},
			{{0, 1, 0, 1}, {0, 1, 0, 2}}},
			"" +
				" _____    _____ \n" +
				"|     |  |     |\n" +
				"|     |  |     |\n" +
				"| [ ] |  | [-] |\n" +
				"|     |  |     |\n" +
				"|_____|  |_____|\n" +
				"\n" +
				" _____    _____ \n" +
				"|     |  |     |\n" +
				"|     |  | [ ] |\n" +
				"| [#] |  |     |\n" +
				"|     |  | [ ] |\n" +
				"|_____|  |_____|\n" +
				"\n" +
				" _____    _____ \n" +
				"|     |  |     |\n" +
				"| [-] |  | [#] |\n" +
				"|     |  |     |\n" +
				"| [-] |  | [#] |\n" +
				"|_____|  |_____|\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := renderer{cards: tt.cards}
			if got := r.getCardString(); got != tt.want {
				t.Errorf("getCardOutputString() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}
