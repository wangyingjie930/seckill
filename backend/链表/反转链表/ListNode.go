package ListNode

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

type List struct {
	size uint64
	head *ListNode
	tail *ListNode
}

func NewList(arr []int) *List {
	list := initList()
	for _, v := range arr {
		list.append(&ListNode{Val: v, Next: nil})
	}
	return list
}

func (list *List) ShowList() {
	node := list.head
	for node != nil {
		fmt.Print(" ", node.Val)
		node = node.Next
	}
}

func initList() *List {
	return &List{
		size: 0,
		head: nil,
		tail: nil,
	}
}

func (list *List) append(node *ListNode) {
	if list.size == 0 {
		list.head = node
		list.tail = node
	}else {
		tailNode := list.tail
		tailNode.Next = node
		list.tail = node
	}
	(*list).size ++
}

/**
反转链表, 参考: https://leetcode-cn.com/problems/reverse-linked-list/solution/dong-hua-yan-shi-206-fan-zhuan-lian-biao-by-user74/
记得反转完之后要重置首尾指针 */
func (list *List) ReverseList() {
	var pre *ListNode
	curr := list.head
	list.tail = list.head

	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}

	list.head = pre
}