package List

/**
使用双指针进行实现
 */
func ReverseList(head *Node) *Node {
	if head == nil {
		return nil
	}
	var pre, cur *Node
	cur = head
	pre = nil
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
