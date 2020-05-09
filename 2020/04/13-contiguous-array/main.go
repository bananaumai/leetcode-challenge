package main

func findMaxLength(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	ys := make(map[int][]int)
	ys[0] = []int{0}

	var y int
	for x, n := range nums {
		if n == 0 {
			y -= 1
		} else {
			y += 1
		}
		ys[y] = append(ys[y], x+1)
	}

	var longest int
	for _, xs := range ys {
		if len(xs) < 2 {
			continue
		}
		distance := xs[len(xs)-1] - xs[0]

		if distance > longest {
			longest = distance
		}
	}

	return longest
}
