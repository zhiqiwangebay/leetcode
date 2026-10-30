package leetcode

func maxArea(height []int) int {

	left, right := 0, len(height)-1
	maxValue := (right - left) * min(height[left], height[right])
	for left < right {
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		if left < right {
			maxValue = max(maxValue, (right-left)*min(height[left], height[right]))
		}
	}
	return maxValue
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
