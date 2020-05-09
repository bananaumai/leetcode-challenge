package main

func findMaxLength(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	var (
		answer int
		count0 int
		count1 int
	)

	for i, n := range nums {
		if n == 0 {
			count0++
		} else {
			count1++
		}

		if count0 == 0 || count1 == 0 {
			continue
		}

		more := max(count0, count1)
		less := min(count0, count1)

		if more*2 > i+1 {
			continue
		}

		answer = max(answer, less*2)
	}

	return answer
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
