package leetcode

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i, j := 0, len(nums)-1
	for i < j {
		for i < j && nums[j] >= nums[0] {
			j--
		}
		for i < j && nums[i] <= nums[0] {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[0], nums[i] = nums[i], nums[0]
	quickSort(nums[:i])
	quickSort(nums[i+1:])
}
