package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	itemsMap map[*ListItem]*ListItem
	len      int
	front    *ListItem
	back     *ListItem
}

func NewList() List {
	newList := new(list)
	newList.itemsMap = make(map[*ListItem]*ListItem)
	return newList
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newListItem := &ListItem{
		Value: v,
		Next:  l.Front(),
		Prev:  nil,
	}

	if l.len == 0 {
		l.back = newListItem
	}
	if l.Front() != nil {
		l.Front().Prev = newListItem
	}
	l.itemsMap[newListItem] = newListItem
	l.front = newListItem
	l.len++

	return newListItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newListItem := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  l.Back(),
	}

	if l.len == 0 {
		l.front = newListItem
	}
	if l.Back() != nil {
		l.Back().Next = newListItem
	}
	l.itemsMap[newListItem] = newListItem
	l.back = newListItem
	l.len++

	return newListItem
}

func (l *list) Remove(i *ListItem) {
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	l.len--
	delete(l.itemsMap, i)
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.Front() {
		return
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	l.Front().Prev = i
	i.Next = l.Front()
	i.Prev = nil
	l.front = i
}
