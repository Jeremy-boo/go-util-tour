package linkedlist

// Object singly linked list value
type Object interface{}

// Node is an element of a single linked list.
type Node struct {
	value Object
	next  *Node
}
