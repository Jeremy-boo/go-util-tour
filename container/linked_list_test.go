package container

import (
	"testing"
)

func TestLinkedList_Len(t *testing.T) {
	l := NewLinkedList()
	l.PushEnd("1111111")
	t.Log(l.Len())
}
