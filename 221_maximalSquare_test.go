package leetcode

import "testing"

func TestMaximalSquare(t *testing.T) {
	tests := []struct {
		name string
		rows []string
		want int
	}{
		{"empty matrix", nil, 0},
		{"empty row", []string{""}, 0},
		{"single zero", []string{"0"}, 0},
		{"single one", []string{"1"}, 1},
		{"single row", []string{"0110"}, 1},
		{"single column", []string{"0", "1", "1"}, 1},
		{"all zeros", []string{"00", "00"}, 0},
		{"first row only", []string{"01", "00"}, 1},
		{"first column only", []string{"00", "10"}, 1},
		{"top left only", []string{"10", "00"}, 1},
		{"limited by left neighbor", []string{"111", "111", "011"}, 4},
		{"limited by upper neighbor", []string{"110", "111", "111"}, 4},
		{"limited by diagonal", []string{"011", "111", "111"}, 4},
		{"full square", []string{"111", "111", "111"}, 9},
		{"rectangular matrix", []string{"10100", "10111", "11111", "10010"}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := make([][]byte, len(tt.rows))
			for row, values := range tt.rows {
				matrix[row] = []byte(values)
			}
			if got := maximalSquare(matrix); got != tt.want {
				t.Errorf("maximalSquare(%q) = %d, want %d", tt.rows, got, tt.want)
			}
		})
	}
}
