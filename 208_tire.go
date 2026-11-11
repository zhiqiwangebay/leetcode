package leetcode

type Trie struct {
	root *Node
}

type Node struct {
	c        rune
	children map[rune]*Node
	isEnd    bool
}

func Constructor() Trie {
	return Trie{root: &Node{children: make(map[rune]*Node)}}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, ch := range word {
		if node.children[ch] == nil {
			node.children[ch] = &Node{c: ch, children: make(map[rune]*Node)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for _, ch := range word {
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, ch := range prefix {
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
