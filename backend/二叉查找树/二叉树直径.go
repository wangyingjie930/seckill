package BST

/**
从二叉树的节点a出发，可以向上或者向下走，但沿途的节点只能经过一次，
到达节点b时路径上的节点个数叫作a到b的距离，那么二叉树任何两个节点之间都有距离，求整棵树上的最大距离。
 */

func DiameterOfBinaryTree(root *Node) int {
	//1. 当x参与的时候: 取 左的子树高度 + 右子树的高度 + 1
	//2. 当x不参与的时候: 取 左子树最长距离 或者 右子树最长距离
	//3种取最大值
	_, distance := processDiameter(root)
	return distance
}

func processDiameter(node *Node) (int, int) {
	if node == nil {
		return 0, 0
	}
	leftHeight, leftDistance := processDiameter(node.left)
	rightHeight, rightDistance := processDiameter(node.right)

	maxDistance := leftHeight + rightHeight + 1
	maxHeight := leftHeight
	if maxDistance < leftDistance {
		maxDistance = leftDistance
	}
	if maxDistance < rightDistance {
		maxDistance = rightDistance
	}
	if maxHeight < rightHeight {
		maxHeight = rightHeight
	}
	return maxHeight + 1, maxDistance
}
