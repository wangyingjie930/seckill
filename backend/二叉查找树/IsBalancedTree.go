package BST

import "math"

func IsBalance(head *Node) (bool, int)  {
	if head == nil {
		return true, 0
	}

	leftBalance, leftHeight := IsBalance(head.left)
	rightBalance, rightHeight := IsBalance(head.right)

	if leftBalance && rightBalance && math.Abs(float64(leftHeight) - float64(rightHeight)) <= 1 {
		return true, int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
	}
	return false, int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}
