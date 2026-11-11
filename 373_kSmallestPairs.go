package leetcode

import "container/heap"

type pair struct {
	i, j int
}

type pairHeap struct {
	pairs        []pair
	nums1, nums2 []int
}

func (h *pairHeap) Len() int {
	return len(h.pairs)
}

func (h *pairHeap) Less(i, j int) bool {
	sumI := h.nums1[h.pairs[i].i] + h.nums2[h.pairs[i].j]
	sumJ := h.nums1[h.pairs[j].i] + h.nums2[h.pairs[j].j]
	return sumI < sumJ
}

func (h *pairHeap) Swap(i, j int) {
	h.pairs[i], h.pairs[j] = h.pairs[j], h.pairs[i]
}

func (h *pairHeap) Push(x interface{}) {
	h.pairs = append(h.pairs, x.(pair))
}

func (h *pairHeap) Pop() interface{} {
	last := len(h.pairs) - 1
	item := h.pairs[last]
	h.pairs = h.pairs[:last]
	return item
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	result := [][]int{}
	if k <= 0 || len(nums1) == 0 || len(nums2) == 0 {
		return result
	}

	minHeap := &pairHeap{
		nums1: nums1,
		nums2: nums2,
	}

	for i := 0; i < len(nums1) && i < k; i++ {
		minHeap.pairs = append(minHeap.pairs, pair{i, 0})
	}
	heap.Init(minHeap)

	for len(result) < k && minHeap.Len() > 0 {
		current := heap.Pop(minHeap).(pair)
		result = append(result, []int{nums1[current.i], nums2[current.j]})

		if current.j+1 < len(nums2) {
			heap.Push(minHeap, pair{current.i, current.j + 1})
		}
	}

	return result
}
