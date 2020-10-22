package linkedlist

import (
	"testing"
)

func TestNew(t *testing.T) {
	l := New()
	t.Log(l)
}

func TestSingleLinkedList_Add(t *testing.T) {
	l := New()
	l.Add([]string{"1111"})
	l.Add([]string{"2222"})
	t.Log(l)
}

func TestSingleLinkedList_IsEmpty(t *testing.T) {
	l := New()
	l.Add(1111)
	if !l.IsEmpty() {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func TestSingleLinkedList_PushEnd(t *testing.T) {
	l := New()
	l.Add("111111")
	l.PushEnd("22222222")
	val, err := l.FindByIndex(1)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if val.(string) != "22222222" {
		t.Error("push end fail")
		return
	}
	t.Log("success")
}

func TestSingleLinkedList_PushFront(t *testing.T) {
	l := New()
	l.Add("111111")
	l.PushFront("22222222")
	val, err := l.FindByIndex(0)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if val.(string) != "22222222" {
		t.Error("push front fail")
		return
	}
	t.Log("success")
}

func TestSingleLinkedList_RemoveFront(t *testing.T) {
	l := New()
	l.Add("111111")
	l.PushFront("22222222")
	l.RemoveFront()
	val, err := l.FindByIndex(0)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if val.(string) != "111111" {
		t.Error("remove front fail")
		return
	}
	t.Log("success")
}

func TestSingleLinkedList_Len(t *testing.T) {
	l := New()
	l.Add("111111")
	l.Add("111111")
	l.Add("111111")
	l.Add("111111")
	len := l.Len()
	if len != 4 {
		t.Error("get length fail")
		return
	}
	t.Log("success")
}

func TestSingleLinkedList_Insert(t *testing.T) {
	l := New()
	l.Add("111111")
	l.Add("111111")
	l.Add("111111")
	l.Add("111111")
	ok, err := l.Insert("222", 2)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if !ok {
		t.Error("add data error")
		return
	}
	val, err := l.FindByIndex(2)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if val.(string) != "222" {
		t.Error("add fail")
		return
	}
	t.Log("success")
}
