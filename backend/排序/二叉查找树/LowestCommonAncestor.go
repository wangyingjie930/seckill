/**
给定两个二叉树的节点node1和node2，找到他们的最低公共祖先节点
 */
package BST

func LowestAncestor(head, o1, o2 *Node) *Node {
	if head == nil || head == o1 || head == o2 {
		return head
	}

	left := LowestAncestor(head.left, o1, o2)
	right := LowestAncestor(head.right, o1, o2)

	if left != nil && right != nil {
		return head
	}else if left == nil {
		return right
	}
	return left
}