package leetcode

import "sort"

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	h := 0
	for _, count := range citations {
		if count < h+1 {
			break
		}
		h++
	}
	return h
}
