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
func buildTree(inorder []int, postorder []int) *TreeNode {
	indexByValue := make(map[int]int, len(inorder))
	for i, value := range inorder {
		indexByValue[value] = i
	}

	postIndex := len(postorder) - 1
	var build func(left, right int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		rootValue := postorder[postIndex]
		postIndex--
		rootIndex := indexByValue[rootValue]

		root := &TreeNode{Val: rootValue}
		root.Right = build(rootIndex+1, right)
		root.Left = build(left, rootIndex-1)
		return root
	}

	return build(0, len(inorder)-1)
}
