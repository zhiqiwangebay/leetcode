package leetcode

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
	if len(grid) == 0 {
		return nil
	}

	var build func(row, col, size int) *Node
	build = func(row, col, size int) *Node {
		if size == 1 {
			return &Node{
				Val:    grid[row][col] == 1,
				IsLeaf: true,
			}
		}
		if size <= 4 {
			sum := 0
			for i := row; i < row+size; i++ {
				for j := col; j < col+size; j++ {
					sum += grid[i][j]
				}
			}
			if sum == 0 || sum == size*size {
				return &Node{
					Val:    sum > 0,
					IsLeaf: true,
				}
			}
		}

		half := size / 2
		topLeft := build(row, col, half)
		topRight := build(row, col+half, half)
		bottomLeft := build(row+half, col, half)
		bottomRight := build(row+half, col+half, half)

		if topLeft.IsLeaf && topRight.IsLeaf &&
			bottomLeft.IsLeaf && bottomRight.IsLeaf &&
			topLeft.Val == topRight.Val &&
			topLeft.Val == bottomLeft.Val &&
			topLeft.Val == bottomRight.Val {
			return &Node{
				Val:    topLeft.Val,
				IsLeaf: true,
			}
		}

		return &Node{
			IsLeaf:      false,
			TopLeft:     topLeft,
			TopRight:    topRight,
			BottomLeft:  bottomLeft,
			BottomRight: bottomRight,
		}
	}

	return build(0, 0, len(grid))
}
