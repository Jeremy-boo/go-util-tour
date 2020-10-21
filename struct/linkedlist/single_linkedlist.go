package linkedlist

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
	Insert(value Object, index int)
	RemoveFront() Node
	RemoveEnd() Node
	Remove(index int) Node
	FindByIndex(index int) Node
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
