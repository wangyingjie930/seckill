package main

import (
	"fmt"
	ListNode "imooc-product/backend/链表/反转链表"
	"math/rand"
)

func main() {
	arr := rand.Perm(50)
	fmt.Println(arr)
	list := ListNode.NewList(arr)
	list.ReverseList()
	list.ShowList()
}
