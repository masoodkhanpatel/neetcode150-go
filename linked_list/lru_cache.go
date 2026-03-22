package linked_list

type LRUNode struct {
	Key, Val   int
	Prev, Next *LRUNode
}

type LRUCache struct {
	Cap        int
	Cache      map[int]*LRUNode
	Left, Right *LRUNode
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		Cap:   capacity,
		Cache: make(map[int]*LRUNode),
		Left:  &LRUNode{Key: 0, Val: 0},
		Right: &LRUNode{Key: 0, Val: 0},
	}
	lru.Left.Next, lru.Right.Prev = lru.Right, lru.Left
	return lru
}

func (this *LRUCache) remove(node *LRUNode) {
	prev, next := node.Prev, node.Next
	prev.Next, next.Prev = next, prev
}

func (this *LRUCache) insert(node *LRUNode) {
	prev, next := this.Right.Prev, this.Right
	prev.Next = node
	next.Prev = node
	node.Prev, node.Next = prev, next
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Cache[key]; ok {
		this.remove(node)
		this.insert(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Cache[key]; ok {
		this.remove(node)
	}
	this.Cache[key] = &LRUNode{Key: key, Val: value}
	this.insert(this.Cache[key])

	if len(this.Cache) > this.Cap {
		lru := this.Left.Next
		this.remove(lru)
		delete(this.Cache, lru.Key)
	}
}
