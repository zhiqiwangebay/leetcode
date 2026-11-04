package leetcode

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	var nHead, nTail *Node
	m := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		n := &Node{Val: cur.Val}
		m[cur] = n
		if nHead == nil {
			nHead = n
			nTail = n
		} else {
			nTail.Next = n
			nTail = n
		}
		cur = cur.Next
	}
	for oldNode, newNode := range m {
		if oldNode.Random != nil {
			newNode.Random = m[oldNode.Random]
		}
	}
	return nHead
}
