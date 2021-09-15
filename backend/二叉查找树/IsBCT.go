package BST

/**
是否为一颗完全二叉搜索树

1, 任意节点有右无左, 返回false
2. 如果遇到第一个只有左没有右的节点, 那么后续的节点都为叶子节点
 */
func IsBCT(head *Node) bool  {
	if head == nil {
		return true
	}
	queue := []*Node{head}
	tmp := false
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.right != nil && node.left == nil {
			return false
		}else if node.left != nil && node.right == nil {
			tmp = true
		}
		if tmp && (node.right != nil || node.left != nil) {
			return false
		}
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
	return true
}
