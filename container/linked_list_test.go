package container

import (
	"github.com/jeremy-boo/go-util-tour/struct/linkedlist"
	"testing"
)

func TestLinkedList_Len(t *testing.T) {
	l := linkedlist.New()
	l.Add(111)
	l.Add(2222)
}
