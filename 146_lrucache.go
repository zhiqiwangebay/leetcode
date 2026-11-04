package leetcode

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	count    int
	m        map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		count:    0,
		m:        make(map[int]*Node),
		head:     &Node{},
		tail:     &Node{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		node.value = value
		this.moveToHead(node)
		return
	}
	newNode := &Node{
		key:   key,
		value: value,
	}
	this.m[key] = newNode
	this.addNode(newNode)
	this.count++
	if this.count > this.capacity {
		tail := this.tail.prev
		this.removeNode(tail)
		delete(this.m, tail.key)
		this.count--
	}

}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addNode(node)
}

func (this *LRUCache) removeNode(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) addNode(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
