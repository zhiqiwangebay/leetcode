package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	vals := make([]int, 0)
	inorderTraversal(root, &vals)
	minDiff := int(^uint(0) >> 1)
	for i := 1; i < len(vals); i++ {
		if vals[i]-vals[i-1] < minDiff {
			minDiff = vals[i] - vals[i-1]
		}
	}
	return minDiff
}

func inorderTraversal(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left, vals)
	*vals = append(*vals, root.Val)
	inorderTraversal(root.Right, vals)
}
