package leetcode

import "testing"

func TestSnakesAndLadders(t *testing.T) {
	tests := []struct {
		name  string
		board [][]int
		want  int
	}{
		{"shortest roll", [][]int{{-1, -1}, {-1, -1}}, 1},
		{"snake back to start", [][]int{{-1, -1}, {-1, 1}}, 1},
		{"unreachable", [][]int{
			{1, -1, -1},
			{1, 1, 1},
			{-1, 1, 1},
		}, -1},
		{"one transfer per roll", [][]int{
			{-1, -1, -1, -1},
			{-1, -1, -1, -1},
			{16, -1, -1, -1},
			{-1, 8, -1, -1},
		}, 2},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := snakesAndLadders(test.board); got != test.want {
				t.Errorf("snakesAndLadders() = %d, want %d", got, test.want)
			}
		})
	}
}
