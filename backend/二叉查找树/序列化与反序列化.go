package BST

import (
	"strconv"
	"strings"
)

func Serialize(head *Node) string  {
	if head == nil {
		return "#,"
	}
	str := strconv.Itoa(head.Key) + ","
	str = str + Serialize(head.left)
	str = str + Serialize(head.right)
	return str
}

func UnSerialize(str string) *Node {
	return reconByString(&str)
}

/**
使用的是队列
 */
func reconByQueue(queue *[]string) *Node {
	val := (*queue)[0]
	*queue = (*queue)[1:]
	if val == "#" {
		return nil
	}
	num, _ := strconv.Atoi(val)
	node := &Node{Key: num, Value: num}
	node.left = reconByQueue(queue)
	node.right = reconByQueue(queue)
	return node
}

/**
使用的是字符串截取
反序列化和序列化调用递归的顺序是一样的
 */
func reconByString(str *string) *Node  {
	index := strings.Index(*str, ",")
	val := (*str)[0:index]
	*str = (*str)[index + 1:]
	if val == "#" {
		return nil
	}
	num, _ := strconv.Atoi(val)
	node := &Node{Key: num, Value: num}
	node.left = reconByString(str)
	node.right = reconByString(str)
	return node
}