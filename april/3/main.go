package main

import "math"

func maxSubArray(nums []int) int {
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums))
	}

	max := math.MinInt64
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if j == i {
				dp[i][j] = nums[j]
			} else {
				dp[i][j] = dp[i][j-1] + nums[j]
			}
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}
	return max
}
