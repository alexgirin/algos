package ds

import (
	"cmp"
	"fmt"
)

type listNode[T cmp.Ordered] struct {
	value T
	next  *listNode[T]
}

type LinkedList[T cmp.Ordered] struct {
	head   *listNode[T]
	length int
}

func NewLinkedList[T cmp.Ordered]() *LinkedList[T] {
	return &LinkedList[T]{
		head:   &listNode[T]{},
		length: 0,
	}
}

func (list *LinkedList[T]) isEmpty() bool {
	return list.head.next == nil
}

func (list *LinkedList[T]) Size() int {
	return list.length
}

func (list *LinkedList[T]) GetFirst() (T, error) {
	if list.isEmpty() {
		return list.head.value, fmt.Errorf("error: %s", "Empty linked list")
	}

	return list.head.next.value, nil
}

func (list *LinkedList[T]) GetLast() (T, error) {
	if list.isEmpty() {
		return list.head.value, fmt.Errorf("error: %s", "Empty linked list")
	}

	cur := list.head
	for cur.next != nil {
		cur = cur.next
	}

	return cur.value, nil
}

func (list *LinkedList[T]) Push(val T) {
	list.length++
	node := &listNode[T]{value: val}

	if list.isEmpty() {
		list.head.next = node
		return
	}

	cur := list.head
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = node
}

func (list *LinkedList[T]) Pop() (T, error) {
	if list.isEmpty() {
		return list.head.value, fmt.Errorf("error: %s", "Empty linked list")
	}

	cur := list.head
	for cur.next.next != nil {
		cur = cur.next
	}
	val := cur.next.value
	cur.next = nil
	list.length--

	return val, nil
}

func (list *LinkedList[T]) Insert(index int, val T) error {
	if index < 0 || index > list.length {
		return fmt.Errorf("error: %s", "Index is out of range")
	}

	node := &listNode[T]{value: val}
	cur := list.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	if cur.next != nil {
		node.next = cur.next
	}

	cur.next = node
	list.length++

	return nil
}

func (list *LinkedList[T]) Enqueue(val T) {
	node := &listNode[T]{value: val}
	list.length++

	if list.isEmpty() {
		list.head.next = node
		return
	}

	node.next = list.head.next
	list.head.next = node
}

func (list *LinkedList[T]) Dequeue() (T, error) {
	if list.isEmpty() {
		return list.head.value, fmt.Errorf("error: %s", "Empty linked list")
	}

	cur := list.head.next
	list.head.next = cur.next
	cur.next = nil
	list.length--

	return cur.value, nil
}

func (list *LinkedList[T]) Delete(index int) (T, error) {
	if index < 0 || index > list.length-1 {
		return list.head.value, fmt.Errorf("error: %s", "Index is out of range")
	}

	cur := list.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	toDelete := cur.next
	cur.next = toDelete.next
	toDelete.next = nil
	list.length--

	return toDelete.value, nil
}

func (list *LinkedList[T]) Contains(val T) bool {
	if list.isEmpty() {
		return false
	}

	cur := list.head
	for cur.next != nil {
		if cur.next.value == val {
			return true
		}
		cur = cur.next
	}

	return false
}

func (list *LinkedList[T]) ToSlice() []T {
	arr := []T{}

	if list.isEmpty() {
		return arr
	}

	cur := list.head
	for cur.next != nil {
		arr = append(arr, cur.next.value)
		cur = cur.next
	}

	return arr
}
