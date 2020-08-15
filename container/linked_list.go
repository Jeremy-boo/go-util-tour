package container

import "errors"

type LinkedList struct {
	head *Node
	size int
}

// IsEmpty
func (list LinkedList) IsEmpty() bool {
	return list.head == nil
}

// Len
func (list LinkedList) Len() int {
	return list.size
}

// NewLinkedList 初始化
func  NewLinkedList() *LinkedList {
	singleLinkedList := new(LinkedList)
	singleLinkedList.size = 0
	singleLinkedList.head = nil
	return singleLinkedList
}

// PushFront 头结点插入元素
func (list LinkedList) PushFront(v interface{}) {
	node := &Node{
		value: v,
		next:  nil,
	}
	if !list.IsEmpty() {
		node.next = list.head
	}
	list.head = node
 	list.size++
}

// PushEnd 尾节点插入元素
func (list *LinkedList) PushEnd(v interface{}) {
	node := &Node{
		value: v,
		next:  nil,
	}
	if list.IsEmpty() {
		list.head = node
	}else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	list.size++
}

// Insert Adds an element to the linked list at the specified location
func (list *LinkedList) Insert(index int,value interface{}) error {
	if index > list.Len() {
		return errors.New("add out of index error")
	}
	if index  == 0 {
		list.PushFront(value)
		return nil
	}
	preNode := list.head
	newNode := &Node{
		value: value,
		next:  nil,
	}
	for i :=1;i < list.Len();i++ {
		if index == i {
			newNode.next = preNode.next
			preNode.next = newNode
			list.size++
			return nil
		}
		preNode = preNode.next
	}
	return errors.New("add failure")
}

func (list *LinkedList) RemoveFront() Object {
	return nil
}

func (list *LinkedList) RemoveEnd() Object {
	return nil
}

func (list *LinkedList) Get(index int) (Object,error) {
	if index > list.Len() {
		return nil,errors.New("index out of list'length")
	}
	return nil,nil
}

func (list *LinkedList) GetHead() *Node {
	return list.head
}