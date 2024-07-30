package algo

import "container/list"





type LRUCache[T any] struct {
	order *list.List
	items map[string]*Node[T]
	capacity int
}

type Node[T any] struct {
	Item T
	ItemPtr *list.Element
}

func NewLRUCache[T any] (capacity int) *LRUCache[T] {
	cache := &LRUCache[T]{}
	cache.order = &list.List{}
	cache.items = make(map[string]*Node[T])
	cache.capacity = capacity
	return cache
}


func (c *LRUCache[T]) get(key string) T {
	var zero T
	if item, ok := c.items[key]; ok {
		zero = item.Item
		c.order.MoveToFront(item.ItemPtr)
	}
	return zero
}

func (c *LRUCache[T]) put(key string, val T) {
	if item, ok := c.items[key]; ok {
		item.Item = val
		c.order.MoveToFront(item.ItemPtr)
		return 
	} 
	// not exist
	if (c.capacity == len(c.items)) {
		back := c.order.Back()
		c.order.Remove(back)
		delete(c.items, back.Value.(string))
		
	}
	c.items[key] = &Node[T]{val, c.order.PushFront(key)}
}