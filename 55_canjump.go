package leetcode

func canJump(nums []int) bool {
	n := len(nums)
	maxReach := 0
	for i := 0; i < n; i++ {
		if i > maxReach {
			return false
		}
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
	}

	return true
}
