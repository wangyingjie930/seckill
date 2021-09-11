package BST

import "fmt"

/**
结点
 */
type Node struct {
	Key int
	Value int
	left *Node
	right *Node
}

type BST struct {
	root *Node
	count int
}

func NewBST() BST {
	return BST{root: nil, count: 0}
}

func (b *BST) IsEmpty() bool {
	return b.count == 0
}

func (b *BST) Size() int {
	return b.count
}

func (b *BST) Insert(key, value int) {
	b.root = b.insert(b.root, key, value)
}

func (b *BST) Contain(key int) bool {
	return b.contain(b.root, key)
}

func (b *BST) Search(key int) *int  {
	return b.search(b.root, key)
}

//前序遍历
func (b *BST) PreOrder()  {
	preOrder(b.root)
}

//中序遍历
func (b *BST) InOrder()  {
	inOrder(b.root)
}

func (b *BST) PostOrder()  {
	postOrder(b.root)
}

func (b *BST) LevelOrder() {

}

func (b *BST) Minimum() int {
	node := minimum(b.root)
	return node.Key
}

func (b *BST) Maximum() int {
	node := maximum(b.root)
	return node.Key
}

func (b *BST) RemoveMin() {
	if b.root != nil {
		removeMin(b.root)
	}
}

func (b *BST) RemoveMax() {
	if b.root != nil {
		removeMax(b.root)
	}
}

func (b *BST) Remove(key int)  {
	b.root = remove(b.root, key)
}

//思考: 二分搜索树的floor和ceil (最小大于58的节点和最大小于58的节点是什么)
//思考: 二分搜索树的rank(58在树中排第几)和select(第10名的节点是什么), 对每个node添加属性: 左右子树加上自身的数量, 保存一份
//思考: 支持重复元素的树, node添加属性: 重复个数

/**
插入节点
 */
func (b *BST) insert (node *Node, key, value int) *Node {
	if node == nil {
		b.count ++
		return &Node{Key: key, Value: value, right: nil, left: nil}
	}

	if node.Key == key {
		node.Value = value
	}else if node.Key > key {
		node.left = b.insert(node.left, key, value)
	}else {
		node.right = b.insert(node.right, key, value)
	}
	return node
}

/**
key是否包含在树中
 */
func (b *BST) contain (node *Node, key int) bool {
	if node == nil {
		return false
	}
	if node.Key == key {
		return true
	}else if node.Key > key {
		return b.contain(node.left, key)
	}else {
		return b.contain(node.right, key)
	}
}

/**
返回查找节点的值
 */
func (b *BST) search (node *Node, key int) *int {
	if node == nil {
		return nil
	}
	if node.Key == key {
		return &node.Value
	}else if node.Key > key {
		return b.search(node.left, key)
	}else {
		return b.search(node.right, key)
	}
}

/**
前序遍历
 */
func preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.Key, " ")

	preOrder(node.left)
	preOrder(node.right)
}

/**
中序遍历
 */
func inOrder(node *Node) {
	if node == nil {
		return
	}
	inOrder(node.left)
	fmt.Print(node.Key, " ")
	inOrder(node.right)
}

/**
后序遍历
 */
func postOrder(node *Node) {
	if node == nil {
		return
	}
	postOrder(node.left)
	postOrder(node.right)
	fmt.Print(node.Key, " ")
}

/**
最小节点, 递归获取左节点就可以找到
 */
func minimum(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return minimum(node.left)
}

/**
最大节点, 递归获取右节点可以找到
 */
func maximum(node *Node) *Node {
	if node.right == nil {
		return node
	}
	return maximum(node.right)
}

/**
删除最大节点
规律: 递归获取右节点可以找到, 而且它最多也只有左子树
 */
func removeMax(node *Node) *Node {
	if node.right == nil {
		left := node.left
		node = nil
		return left
	}
	node.right = removeMax(node.right)
	return node
}

/**
删除最小节点
规律: 递归获取左节点就可以找到, 而且它做多只有右子树
 */
func removeMin(node *Node) *Node  {
	if node.left == nil {
		right := node.right
		node = nil
		return right
	}
	node.left = removeMin(node.left)
	return node
}

/**
删除节点
 */
func remove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if key < node.Key {
		//递归向左找
		node.left = remove(node.left, key)
		return node
	}else if key > node.Key {
		//向右找
		node.right = remove(node.right, key)
		return node
	}else {
		if node.right == nil {
			//只含有左子树
			leftNode := node.left
			node = nil
			return leftNode
		}else if node.left == nil {
			//只含有右子树
			rightNode := node.right
			node = nil
			return rightNode
		}else {
			//同时又左右子树
			//找出后继节点作为新的节点(被删除节点的右子树的最小节点)
			successor := minimum(node.right)
			//重新赋值新节点的左右节点(被删除节点的左右节点)
			successor.right = removeMin(node.right)
			successor.left = node.left
			node = nil
			return successor
		}
	}
}

