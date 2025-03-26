package linked

type LRUCache struct {
	m        map[int]*Node
	head     *Node
	tail     *Node
	capacity int
}

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:        make(map[int]*Node),
		capacity: capacity,
	}
}

func (cache *LRUCache) Get(key int) int {
	if v := cache.get(key); v != nil {
		return v.value
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if v := cache.get(key); v != nil {
		v.value = value
		return
	}

	newNode := &Node{
		key:   key,
		value: value,
	}
	cache.m[key] = newNode

	if cache.head == nil {
		cache.head = newNode
		cache.tail = newNode
		return
	}

	newNode.next = cache.head
	cache.head.prev = newNode
	cache.head = newNode

	if len(cache.m) > cache.capacity {
		delete(cache.m, cache.tail.key)
		cache.tail = cache.tail.prev
		cache.tail.next = nil
	}
}

func (cache *LRUCache) get(key int) *Node {
	v, ok := cache.m[key]

	if !ok {
		return nil
	}

	if v == cache.head {
		return v
	}

	prev := v.prev
	next := v.next

	if v == cache.tail {
		cache.tail = prev
	}

	cache.head.prev = v
	prev.next = next
	if next != nil {
		next.prev = prev
	}

	v.prev = nil
	v.next = cache.head
	cache.head = v
	return v
}
