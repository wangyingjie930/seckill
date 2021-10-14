package practice

import (
	"strconv"
)

func NumDecodings(s string) int {
	bytes := []byte(s)
	if s == "0" || (len(s) > 0 && string(bytes[0]) == "0") {
		return 0
	}
	if len(s) <= 1 {
		return 1
	}
	num := 0
	if len(s) > 1 {
		num = NumDecodings(string(bytes[1:]))

		two, _ := strconv.Atoi(string(bytes[:2]))
		if two <= 26 {
			num = num + NumDecodings(string(bytes[2:]))
		}
	}
	return num
}

func NumDecodings1(s string) int {
	return processNumDecodings(s, 0)
}

/**
使用动态规划的方法进行优化
 */
func processNumDecodings(s string, i int) int {
	if s == "0" || (len(s) > 0 && string(s[0]) == "0") {
		return 0
	}
	dp := make([]int, len(s) + 1)
	dp[len(s)] = 1
	if string(s[len(s) - 1]) != "0" {
		dp[len(s) - 1] = 1
	}
	for i := len(s) - 2; i >= 0; i-- {
		if string(s[i]) == "0" {
			continue
		}
		dp[i] = dp[i + 1]
		if two, err := strconv.Atoi(s[i : i+2]); err == nil && two <= 26 {
			dp[i] = dp[i] + dp[i + 2]
		}
	}
	return dp[i]
}