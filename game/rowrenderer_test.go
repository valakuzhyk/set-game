package game

import "testing"

func Test_getCardLine(t *testing.T) {
	tests := []struct {
		name string
		card *Card
		want string
	}{
		{"no shade, char 0", &Card{shading: 0, shape: 0}, "O  O"},
		{"partial shade, char 1", &Card{shading: 1, shape: 1}, "+--+"},
		{"full shade, char 2", &Card{shading: 2, shape: 2}, "#==#"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := rowRenderer{}
			if got := r.getCardLine(tt.card); got != tt.want {
				t.Errorf("getCardLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
