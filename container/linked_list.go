package container

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

// IsEmpty
func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}

// Len
func (list *LinkedList) Len() int {
	return list.size
}

// NewLinkedList 初始化
func (list *LinkedList) NewLinkedList() {
	list.size = 0
	list.head = nil
	list.tail = nil
}

// PushFront 头结点插入元素
func (list *LinkedList) PushFront(v interface{}) {
	node := &Node{
		value: v,
		next:  nil,
	}
	if list.Len() == 0 {
		list.tail = node
	}else {
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
	if list.Len() == 0 {
		list.head = node
	}else  {
		list.tail.next = node
	}
	list.tail = node
	list.size++
}