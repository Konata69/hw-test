package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type KeyValue struct {
	Key   Key
	Value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	_, ok := c.items[key]
	item := KeyValue{
		Key:   key,
		Value: value,
	}
	if ok {
		c.queue.MoveToFront(c.items[key])
		c.items[key].Value = item
		return ok
	}

	newItem := c.queue.PushFront(item)
	c.items[key] = newItem

	if c.queue.Len() > c.capacity {
		back := c.queue.Back().Value.(KeyValue)
		delete(c.items, back.Key)
		c.queue.Remove(c.queue.Back())
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	_, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(c.items[key])
		value := c.items[key].Value
		if value != nil {
			return value.(KeyValue).Value, ok
		}
		return value, ok
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
