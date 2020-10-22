package main

import (
	"fmt"
	v1 "github.com/go-util-tour/container"
)

func main() {
	linkList := v1.NewLinkedList()
	linkList.PushFront("first")
	linkList.PushEnd("two")
	val, _ := linkList.Get(1)
	fmt.Println(val)
}
