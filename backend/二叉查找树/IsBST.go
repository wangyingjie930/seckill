package BST

import "math"

/**
使用递归实现的, 但是空间复杂度不太行..
 */
func IsValidBST(head *Node) (bool, *Node, *Node) {
	if head == nil {
		return true, nil, nil
	}
	leftIsBST, leftMax, leftMin := IsValidBST(head.left)
	rightIsBST, rightMax, rightMin := IsValidBST(head.right)

	if leftIsBST && rightIsBST {
		if leftMax == nil && rightMin == nil {
			return true, head, head
		}else if leftMax != nil && rightMin == nil && leftMax.Key < head.Key {
			return true, head, leftMin
		}else if rightMin != nil && leftMax == nil && rightMin.Key > head.Key {
			return true, rightMax, head
		}else if leftMax != nil && rightMin != nil && leftMax.Key < head.Key && rightMin.Key > head.Key {
			return true, rightMax, leftMin
		}
	}
	return false, nil, nil
}

func IsValidBST2(head *Node) bool {
	if head == nil {
		return true
	}
	var stack []*Node
	tmp := math.MinInt64
	for len(stack) > 0 || head != nil {
		if head != nil {
			stack = append(stack, head)
			head = head.left
		}else {
			head = stack[len(stack) - 1]
			stack = stack[:(len(stack) - 1)]
			if tmp > head.Key {
				return false
			}else {
				tmp = head.Key
			}
			head = head.right
		}
	}
	return true
}
