package leetcode

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		{"", "", 0},
		{"abc", "", 0},
		{"", "a", -1},
		{"sadbutsad", "sad", 0},
		{"hello", "ll", 2},
		{"mississippi", "issip", 4},
		{"ababababac", "ababac", 4},
		{"aaaaa", "bba", -1},
		{"abc", "abcd", -1},
		{"abc", "abc", 0},
	}

	for _, tt := range tests {
		if got := strStr(tt.haystack, tt.needle); got != tt.want {
			t.Errorf("strStr(%q, %q) = %d, want %d", tt.haystack, tt.needle, got, tt.want)
		}
	}
}
