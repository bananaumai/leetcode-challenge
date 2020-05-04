package main

import "math"

func maxSubArray(nums []int) int {
	m := math.MinInt64
	for i := 0; i < len(nums); i++ {
		for k := i; k < len(nums); k++ {
			m = max(m, sum(nums[i:k+1]))
		}
	}
	return m
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func sum(nums []int) int {
	var acc int
	for _, n := range nums {
		acc += n
	}
	return acc
}
