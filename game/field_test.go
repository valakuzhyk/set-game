package game

import (
	"reflect"
	"testing"
)

func TestField_RedistributeCards(t *testing.T) {
	tests := []struct {
		name string
		f    *Field
		want *Field
	}{
		{"basic",
			&Field{
				{{0, 0, 0, 0}, {0, 0, 0, 1}},
				{{0, 0, 1, 0}, {0, 0, 1, 1}},
				{{0, 0, 2, 0}, {0, 0, 2, 1}},
			},
			&Field{
				{{0, 0, 0, 0}, {0, 0, 0, 1}},
				{{0, 0, 1, 0}, {0, 0, 1, 1}},
				{{0, 0, 2, 0}, {0, 0, 2, 1}},
			},
		},
		{"Move cards",
			&Field{
				{{0, 0, 0, 0}, {0, 0, 0, 1}},
				{{0, 0, 1, 0}},
				{{0, 0, 1, 1}, {0, 0, 2, 0}, {0, 0, 2, 1}},
			},
			&Field{
				{{0, 0, 0, 0}, {0, 0, 0, 1}},
				{{0, 0, 1, 0}, {0, 0, 1, 1}},
				{{0, 0, 2, 0}, {0, 0, 2, 1}},
			},
		},
		{"Less Cards",
			&Field{
				{{0, 0, 0, 0}},
				{},
				{},
			},
			&Field{
				{{0, 0, 0, 0}},
				nil,
				nil,
			},
		},
		{"Remove nil cards",
			&Field{
				{{0, 0, 0, 0}, {0, 0, 0, 1}},
				{{0, 0, 1, 0}},
				{nil, {0, 0, 2, 0}, {0, 0, 2, 1}},
			},
			&Field{
				{{0, 0, 0, 0}, {0, 0, 0, 1}},
				{{0, 0, 1, 0}, {0, 0, 2, 0}},
				{{0, 0, 2, 1}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f.RedistributeCards()
			if !reflect.DeepEqual(tt.f, tt.want) {
				t.Fatalf("got: \n%v want: \n%v ", tt.f, tt.want)
			}
		})
	}
}
