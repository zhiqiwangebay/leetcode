package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		words []string
		want  []int
	}{
		{"example", "barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"repeated word", "wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{8}},
		{"overlapping matches", "barfoofoobarthefoobarman", []string{"bar", "foo", "the"}, []int{6, 9, 12}},
		{"every offset", "aaaaaaaa", []string{"aa", "aa", "aa"}, []int{0, 1, 2}},
		{"nonzero offset", "xbarfoo", []string{"bar", "foo"}, []int{1}},
		{"no match", "aaaaaa", []string{"aa", "bb"}, []int{}},
		{"short string", "foo", []string{"foo", "bar"}, []int{}},
		{"empty words", "foo", []string{}, []int{}},
		{"empty word", "foo", []string{""}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSubstring(tt.s, tt.words)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring(%q, %q) = %v, want %v", tt.s, tt.words, got, tt.want)
			}
		})
	}
}
