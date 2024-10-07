package collections

import "container/list"


type LRUCache[T any] struct {
    Capacity int
    Data     map[string]*Node[T]
    Queue    *list.List
}

type Node[T any] struct {
    Key   string
    Value T
    KeyPtr *list.Element
}



func (c *LRUCache[T]) Put(key string, value T) {
	if item, ok := c.Data[key]; ok {
		item.Value = value
		c.Data[key] = item
		c.Queue.MoveToFront(item.KeyPtr)
		return
	} 
	if c.Queue.Len() >= c.Capacity {
		last := c.Queue.Back()
		c.Queue.Remove(last)
		delete(c.Data, last.Value.(*Node[T]).Key)
	}
	node := &Node[T]{Key: key, Value: value}
	ptr := c.Queue.PushFront(node)
	node.KeyPtr = ptr
	c.Data[key] = node
}


func (c *LRUCache[T]) Get(key string, value T) {

}

