package leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
func connect(root *Node) *Node {
    for level := root; level != nil; {
        dummy := &Node{}
        tail := dummy

        for node := level; node != nil; node = node.Next {
            if node.Left != nil {
                tail.Next = node.Left
                tail = tail.Next
            }
            if node.Right != nil {
                tail.Next = node.Right
                tail = tail.Next
            }
        }

        tail.Next = nil
        level = dummy.Next
    }

    return root
}
