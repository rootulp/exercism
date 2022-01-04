package linkedlist

import (
	"errors"
	"fmt"
)

// Define the List and Element types here.
type List struct {
	head *Element
	size int
}

type Element struct {
	data int
	next *Element
}

func (l *List) String() string {
	return fmt.Sprintf("(size: %d, head: %v)", l.size, l.head)
}

func (e *Element) String() string {
	return fmt.Sprintf("(data: %d, next: %v)", e.data, e.next)
}

func New(array []int) (result *List) {
	result = &List{}
	if array == nil {
		return result
	}
	for _, v := range array {
		result.Push(v)
	}
	return result
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(data int) {
	element := &Element{
		data: data,
	}
	if l.head == nil {
		l.head = element
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = element
	}
	l.size += 1
}

func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("attempted to pop from empty list")
	}
	current := l.head
	for current.next != nil && current.next.next != nil {
		current = current.next
	}
	fmt.Printf("l %v and current %v\n", l, current)
	result := current.next.data
	current.next = nil
	l.size -= 1
	return result, nil
}

func (l *List) Array() (result []int) {
	if l.head == nil {
		return []int{}
	}
	current := l.head
	for current.next != nil {
		result = append(result, current.data)
		current = current.next
	}
	result = append(result, current.data)
	return result
}

func (l *List) Reverse() *List {
	slice := l.Array()
	return New(reverse(slice))
}

func reverse(slice []int) (reversed []int) {
	for _, v := range slice {
		reversed = append([]int{v}, reversed...)
	}
	return reversed
}
