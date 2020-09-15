package main

import (
	"fmt"
	"github.com/go-util-tour/container"
)

func main() {
	linkList := container.NewLinkedList()
	linkList.PushFront("first")
	linkList.PushEnd("two")
	val, _ := linkList.Get(1)
	fmt.Println(val)
}
