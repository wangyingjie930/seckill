package BST

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
