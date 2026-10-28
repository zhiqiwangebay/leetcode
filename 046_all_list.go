package leetcode

func permute(nums []int) (ann [][]int) {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{nums}
	}
	rest := make([]int, 0, len(nums)-1)
	rest = append(rest, nums[1:]...)
	restPerms := permute(rest)
	for _, perm := range restPerms {
		for i := 0; i <= len(perm); i++ {
			newPerm := make([]int, len(perm)+1)
			copy(newPerm[:i], perm[:i])
			newPerm[i] = nums[0]
			copy(newPerm[i+1:], perm[i:])
			ann = append(ann, newPerm)
		}
	}
	return ann
}
