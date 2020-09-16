package container

// LinkedListDouble 双向链表
type LinkedListDouble struct {
	size int
	head *DNode
	tail *DNode
}

// NewLinkedListDouble 初始化
func NewLinkedListDouble() *LinkedListDouble {
	doubleList := new(LinkedListDouble)
	doubleList.size = 0
	doubleList.head = nil
	doubleList.tail = nil
	return doubleList
}

// Len 获取链表长度
func (list *LinkedListDouble) Len() int {
	return list.size
}

func (list *LinkedListDouble) IsEmpty() bool {
	return list.head == nil && list.tail == nil
}

// PushFront 链表头部插入数据
func (list *LinkedListDouble) PushFront(value interface{}) {

}
