package practice

import (
	"fmt"
	"testing"
)

var arr = []int{92,4,2,49,47,32,90,95,35,85,41,84,82,48,6,94,69,20,50,87,79,28,
	67,15,36,55,75,93,91,56,8,12,42,16,66,59,57,5,33,1,54,97,80,77,65,52,29,81,14,13,
	86,64,98,68,46,58,76,70,11,96,60,31,34,44,30,22,61,24,43,26,62,21,74,83,39,72,73,88,45,27,0,10,89,
	51,18,37,17,3,40,9,19,63,71,38,53,78,7,23,99,25}

func TestMajorityElement(t *testing.T) {
	arr := []int{1, 2, 3, 2, 2, 2, 5, 4, 2, 5, 5, 5, 5, 5, 6, 5, 8}
	fmt.Println(MajorityElement(arr))
	//fmt.Print(arr[0:1], arr[1:])
}

func TestFindMedianSortedArrays(t *testing.T) {
	fmt.Println(FindMedianSortedArrays([]int{}, []int{2, 4}))
}
