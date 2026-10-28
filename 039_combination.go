package main

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	path := []int{}

	var search func(start, remaining int)
	search = func(start, remaining int) {
		if remaining == 0 {
			result = append(result, append([]int{}, path...))
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > remaining {
				continue
			}
			path = append(path, candidates[i])
			search(i, remaining-candidates[i])
			path = path[:len(path)-1]
		}
	}

	search(0, target)
	return result
}
