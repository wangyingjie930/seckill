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

func preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.Key, " ")

	preOrder(node.left)
	preOrder(node.right)
}

func inOrder(node *Node) {
	if node == nil {
		return
	}
	inOrder(node.left)
	fmt.Print(node.Key, " ")
	inOrder(node.right)
}

func postOrder(node *Node) {
	if node == nil {
		return
	}
	postOrder(node.left)
	postOrder(node.right)
	fmt.Print(node.Key, " ")
}


