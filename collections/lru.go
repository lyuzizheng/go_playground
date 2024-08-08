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
	} else {

	}

}


func (c *LRUCache[T]) Get(key string, value T) {

}

