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
func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		levelSize := len(q)
		levelSum := 0
		for i := 0; i < levelSize; i++ {
			node := q[i]
			levelSum += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, float64(levelSum)/float64(levelSize))
		q = q[levelSize:]
	}

	return res
}
