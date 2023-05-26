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

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	_, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(c.items[key])
		c.items[key].Value = value
		return ok
	}

	newItem := c.queue.PushFront(value)
	c.items[key] = newItem

	if c.queue.Len() > c.capacity {
		for itemsKey, itemsValue := range c.items {
			if itemsValue == c.queue.Back() {
				delete(c.items, itemsKey)
			}
		}
		c.queue.Remove(c.queue.Back())
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	_, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(c.items[key])
		return c.items[key].Value, ok
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
