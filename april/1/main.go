package main

func singleNumber(nums []int) int {
	var acc int
	for _, n := range nums {
		acc = acc ^ n
	}
	return acc
}
