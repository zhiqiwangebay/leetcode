package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"nil", nil, nil},
		{"empty", []int{}, []int{}},
		{"single", []int{5}, []int{5}},
		{"sorted", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"reversed", []int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
		{"duplicates", []int{3, 1, 3, 2, 1, 3}, []int{1, 1, 2, 3, 3, 3}},
		{"all equal", []int{2, 2, 2, 2}, []int{2, 2, 2, 2}},
		{"negative and positive", []int{0, -4, 7, -1, 0}, []int{-4, -1, 0, 0, 7}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			quickSort(test.nums)
			if !reflect.DeepEqual(test.nums, test.want) {
				t.Errorf("quickSort() = %v, want %v", test.nums, test.want)
			}
		})
	}
}

func TestQuickSortSmallArrays(t *testing.T) {
	for length := 0; length <= 6; length++ {
		count := 1
		for i := 0; i < length; i++ {
			count *= 3
		}
		for value := 0; value < count; value++ {
			nums := make([]int, length)
			code := value
			for i := range nums {
				nums[i] = code%3 - 1
				code /= 3
			}
			want := append([]int{}, nums...)
			sort.Ints(want)
			quickSort(nums)
			if !reflect.DeepEqual(nums, want) {
				t.Fatalf("quickSort() = %v, want %v", nums, want)
			}
		}
	}
}
