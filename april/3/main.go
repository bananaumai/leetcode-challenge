package main

func maxSubArray(nums []int) int {
	sub := subArr(nums)

	var m int
	for _, ns := range sub {
		m = max(m, sum(ns))
	}

	return m
}

func subArr(nums []int) [][]int {
	var sub [][]int
	for i := 0; i < len(nums); i++ {
		for k := i; k < len(nums); k++ {
			sub = append(sub, nums[i:k+1])
		}
	}
	return sub
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
