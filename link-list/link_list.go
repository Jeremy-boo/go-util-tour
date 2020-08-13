package link_list


type Node struct {
	Data interface{}
	Next *Node
}

type LinkList struct {
	head *Node
}

// IsEmpty
func (list *LinkList) IsEmpty() bool {
	return list.head == nil
}