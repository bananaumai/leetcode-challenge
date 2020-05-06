package main

func majorityElement(nums []int) int {
	threshold := len(nums) / 2

	nMap := make(map[int]int)
	for _, n := range nums {
		nMap[n]++
		if nMap[n] > threshold {
			return n
		}
	}

	panic("majority element not found")
}
