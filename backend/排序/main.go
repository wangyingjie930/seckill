package main

import (
	"fmt"
	MergeSort "imooc-product/backend/排序/归并排序"
)

func main() {
	arr := []int{1, 9, 9, 6, 1, 0, 3, 0}
	MergeSort.MergeSort(arr, 0, len(arr) - 1)
	fmt.Print(arr)
}
