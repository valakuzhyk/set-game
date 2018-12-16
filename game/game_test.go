package game

import (
	"reflect"
	"testing"
)

func TestState_drawNewCard(t *testing.T) {
	type fields struct {
		field    [][]*Card
		deck     []*Card
		selected []int
		score    int
	}
	type args struct {
		indices []int
	}
	type want struct {
		field [][]*Card
		deck  []*Card
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			"Basic",
			fields{
				field: [][]*Card{
					{{0, 0, 0, 0}, {0, 0, 0, 1}, {0, 0, 0, 2}},
					{{1, 0, 0, 0}, {1, 0, 0, 1}, {1, 0, 0, 2}},
					{{2, 0, 0, 0}, {2, 0, 0, 1}, {2, 0, 0, 2}},
				},
				deck: []*Card{{1, 1, 1, 1}, {2, 2, 2, 2}},
			},
			args{[]int{1, 8}},
			want{
				field: [][]*Card{
					{{0, 0, 0, 0}, {0, 0, 0, 1}, {0, 0, 0, 2}},
					{{1, 1, 1, 1}, {1, 0, 0, 1}, {1, 0, 0, 2}},
					{{2, 0, 0, 0}, {2, 0, 0, 1}, {2, 2, 2, 2}},
				},
				deck: []*Card{},
			},
		},
		{
			"No cards left",
			fields{
				field: [][]*Card{
					{{0, 0, 0, 0}, {0, 0, 0, 1}, {0, 0, 0, 2}},
					{{1, 0, 0, 0}, {1, 0, 0, 1}, {1, 0, 0, 2}},
					{{2, 0, 0, 0}, {2, 0, 0, 1}, {2, 0, 0, 2}},
				},
				deck: []*Card{{1, 1, 1, 1}, {2, 2, 2, 2}},
			},
			args{[]int{1, 8, 3}},
			want{
				field: [][]*Card{
					{{0, 0, 0, 0}, {0, 0, 0, 2}},
					{{1, 1, 1, 1}, {1, 0, 0, 1}, {1, 0, 0, 2}},
					{{2, 0, 0, 0}, {2, 0, 0, 1}, {2, 2, 2, 2}},
				},
				deck: []*Card{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &State{
				field:    tt.fields.field,
				deck:     tt.fields.deck,
				selected: tt.fields.selected,
				score:    tt.fields.score,
			}
			s.drawNewCard(tt.args.indices...)

			if !reflect.DeepEqual(tt.want.field, s.field) {
				t.Fatalf("Fields not equal")
			}

			if !reflect.DeepEqual(tt.want.deck, s.deck) {
				t.Fatalf("Decks not equal")
			}
		})
	}
}
