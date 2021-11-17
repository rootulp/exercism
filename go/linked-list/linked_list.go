package linkedlist

import "errors"

// Define List and Node types here.
type List struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}

var ErrEmptyList = errors.New("empty list")

func NewList(args ...interface{}) *List {
	list := &List{}
	for _, v := range args {
		list.PushBack(v)
	}
	return list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) PushFront(v interface{}) {
	new := Node{Val: v, next: l.head}

	if l.size == 0 {
		l.head = &new
		l.tail = &new
	} else {
		l.head.prev = &new
		l.head = &new
	}
	l.size++
}

func (l *List) PushBack(v interface{}) {
	new := Node{Val: v, prev: l.tail}

	if l.size == 0 {
		l.head = &new
		l.tail = &new
	} else {
		l.tail.next = &new
		l.tail = &new
	}
	l.size++
}

func (l *List) PopFront() (interface{}, error) {
	if l.size == 0 {
		return nil, ErrEmptyList
	}
	if l.size == 1 {
		result := l.head.Val
		l.head = nil
		l.tail = nil
		l.size -= 1
		return result, nil
	}
	result := l.head.Val
	l.head = l.head.next
	l.head.prev = nil
	l.size -= 1
	return result, nil
}

func (l *List) PopBack() (interface{}, error) {
	if l.size == 0 {
		return nil, ErrEmptyList
	}

	if l.size == 1 {
		result := l.tail.Val
		l.head = nil
		l.tail = nil
		l.size -= 1
		return result, nil
	}
	result := l.tail.Val
	l.tail = l.tail.prev
	l.tail.next = nil
	l.size -= 1
	return result, nil
}

func (l *List) Reverse() {
	data := []interface{}{}
	d, err := l.PopBack()
	for err != ErrEmptyList {
		data = append(data, d)
		d, err = l.PopBack()
	}
	newList := NewList(data...)
	l.head = newList.head
	l.tail = newList.tail
	l.size = newList.size
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
