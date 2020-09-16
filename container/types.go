package container

type Object interface{}

type Node struct {
	value Object
	next  *Node
}

type DNode struct {
	value Object
	prev  *DNode
	next  *DNode
}

// SingleLinkedList is single linked list for The data structure
type SingleLinkedList interface {
	Len() int
	IsEmpty() bool
	NewLinkedList()
	PushFront(value interface{})
	PushEnd(value interface{})
	Insert(index int, value interface{}) error
	RemoveFront() Object
	RemoveEnd() Object
	Get(index int) (Object, error)
	GetHead() *Node
}

type DoubleLinkedList interface {
	Len() int
	PushFront(value interface{})
	PushEnd(value interface{})
}
