package linkedlist

import "errors"

// Object singly linked list value
type Object interface{}

// Node is an element of a single linked list.
type Node struct {
	value Object
	next  *Node
}

// LinkedList LinkedList method
type LinkedList interface {
	Len() int
	IsEmpty() bool
	PushFront(value Object)
	PushEnd(value Object)
	Add(value Object)
	Insert(value Object, index int) (bool, error)
	RemoveFront() bool
	RemoveEnd() bool
	RemoveAll() bool
	Remove(index int) (bool, error)
	FindByIndex(index int) (Object, error)
}

// SingleLinkedList Single list
type SingleLinkedList struct {
	head *Node
	tail *Node
	len  int
}

// New  returns an initialized list.
func New() *SingleLinkedList {
	return new(SingleLinkedList).Init()
}

// Init initializes or clears SingleLinkedList list.
func (list *SingleLinkedList) Init() *SingleLinkedList {
	list.head = nil
	list.tail = nil
	list.len = 0
	return list
}

// IsEmpty Determine whether the linked list is empty.
func (list *SingleLinkedList) IsEmpty() bool {
	return list.head == nil && list.tail == nil
}

// Len returns list's length.
func (list *SingleLinkedList) Len() int {
	return list.len
}

// Add add a value to linked list
func (list *SingleLinkedList) Add(value Object) {
	node := &Node{
		value: value,
		next:  nil,
	}
	if list.head == nil {
		list.head = node
		list.tail = list.head
	} else {
		prev := list.tail
		prev.next = node
		list.tail = node
	}
	list.len++
}

// PushFront Insert data from the head node of the singly linked list.
func (list *SingleLinkedList) PushFront(value Object) {
	node := &Node{
		value: value,
		next:  nil,
	}
	node.next = list.head
	list.head = node
	list.len++
}

// PushEnd Insert data from the tail node of the singly linked list.
func (list *SingleLinkedList) PushEnd(value Object) {
	node := &Node{
		value: value,
		next:  nil,
	}
	prev := list.tail
	prev.next = node
	list.tail = node
	list.len++
}

// Insert Insert data at the specified position in the linked list.
func (list *SingleLinkedList) Insert(value Object, index int) (bool, error) {
	if index > list.Len()-1 || index < 0 {
		return false, errors.New("index is not available")
	}
	// add to head
	if index == 0 {
		list.PushFront(value)
	}
	prev := list.head
	node := &Node{
		value: value,
		next:  nil,
	}
	for i := 1; i < list.Len(); i++ {
		if index == i {
			node.next = prev.next
			prev.next = node
			list.len++
			break
		}
		prev = prev.next
	}
	return true, nil
}

// RemoveFront Delete head node.
func (list *SingleLinkedList) RemoveFront() bool {
	if list.head == nil {
		return false
	}
	list.head = list.head.next
	list.len--
	return true
}

// RemoveEnd Delete tail node.
func (list *SingleLinkedList) RemoveEnd() bool {
	if list.tail == nil {
		return false
	}
	if list.Len() == 1 {
		list.RemoveFront()
	}
	prev := list.head
	for i := 0; i < list.Len(); i++ {
		if i == list.Len()-2 {
			prev.next = nil
			list.tail = prev
			list.len--
			break
		}
		prev = prev.next
	}
	return true
}

// Remove Delete the node at the specified position.
func (list *SingleLinkedList) Remove(index int) (bool, error) {
	if index < 0 || index > list.Len()-1 {
		return false, errors.New("index is not available")
	}
	if index == 0 {
		list.RemoveFront()
	}
	if index == list.Len()-1 {
		list.RemoveEnd()
	}
	prev := list.head
	for i := 1; i < list.Len()-1; i++ {
		if index == i {
			prev.next = prev.next.next
			list.len--
			break
		}
		prev = prev.next
	}
	return true, nil
}

// FindByIndex Find the value of an element based on the index.
func (list *SingleLinkedList) FindByIndex(index int) (Object, error) {
	if index < 0 || index > list.Len()-1 {
		return nil, errors.New("index is not available")
	}
	var val Object
	prev := list.head
	for i := 0; i < list.Len(); i++ {
		if index == i {
			val = prev.value
			break
		}
		prev = prev.next
	}
	return val, nil
}
