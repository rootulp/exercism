package linkedlist

// Define List and Node types here.
type List struct {
	head *Node
	tail *Node
}

type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}

var ErrEmptyList error

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
	panic("Please implement the PushFront function")
}

func (l *List) PushBack(v interface{}) {
	new := Node{Val: v}
	if l.head == nil {
		l.head = &new
	}
	if l.tail == nil {
		l.tail = &new
	} else {
		l.tail.next = &new
		new.prev = l.tail
		l.tail = &new
	}
}

func (l *List) PopFront() (interface{}, error) {
	panic("Please implement the PopFront function")
}

func (l *List) PopBack() (interface{}, error) {
	panic("Please implement the PopBack function")
}

func (l *List) Reverse() {
	panic("Please implement the Reverse function")
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
