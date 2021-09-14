package List

import "fmt"

type Node struct {
	Val int
	Next *Node
}

func NewList(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	head := &Node{arr[0], nil}
	p := head
	for _, v := range arr[1:] {
		node := &Node{v, nil}
		p.Next = node
		p = p.Next
	}
	return head
}

func ShowList(head *Node) {
	node := head
	for node != nil {
		fmt.Print(node.Val, ", ")
		node = node.Next
	}
	fmt.Println()
}


