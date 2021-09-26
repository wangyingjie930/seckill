package practice

/**
定义：数组中累积和与最小值的乘积，假设叫做指标A。给定一个数组，请返回子数组中，指标A最大的值。
 */
//实际就是以当前的数字为最小值, 找左边离他最近的比它小的数, 和右边最近比它小的数

//[5 3 2 1 6 7 8 4]
func QuotaA(nums []int) (int, map[string]int) {
	var stack []int //构建一个由小到大的单调栈
	max := -1
	maxMap := make(map[string]int)
	for i := 0; i < len(nums); i ++ {
		//遇到不符合大小的弹出栈, 直至符合条件
		for len(stack) != 0 && nums[stack[len(stack) - 1]] > nums[i] {
			index := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			end := i - 1
			start := 0
			if len(stack) != 0 {
				start = stack[len(stack) - 1] + 1
			}

			sum := 0
			for j := start; j <= end; j ++ {
				sum = sum + nums[j]
			}
			if max < sum * nums[index] {
				max = sum * nums[index]
				maxMap["start"] = start
				maxMap["end"] = end
			}
		}
		//压入栈
		stack = append(stack, i)
	}

	//清算阶段
	for i := len(stack) - 1; i >= 0; i -- {
		start := 0
		if i >= 1 {
			start = stack[i - 1] + 1
		}
		end := len(stack) - 1

		sum := 0
		for j := start; j <= end; j ++ {
			sum = sum + nums[j]
		}
		if max < sum * nums[i] {
			max = sum * nums[i]
			maxMap["start"] = start
			maxMap["end"] = end
		}
	}
	return max, maxMap
}

