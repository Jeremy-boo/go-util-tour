package container

type Object interface {}


type Node struct {
	value Object
	next *Node
}


// SingleLinkedList is single linked list for The data structure
type SingleLinkedList interface {
	Len() int
	IsEmpty() bool
	NewLinkedList()
	PushFront(value interface{})
	PushEnd(value interface{})
	Insert(index int,value interface{}) error
	RemoveFront() Object
	RemoveEnd() Object
	Get(index int) Object
	GetHead() *Node
}
