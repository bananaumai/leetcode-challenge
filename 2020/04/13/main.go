package main

func findMaxLength(nums []int) int {
	var (
		maxSize int
		size0   int
		size1   int
	)

	for i := 0; i < len(nums); i++ {
		if maxSize >= len(nums)-i {
			break
		}

		for j := i; j < len(nums); j++ {
			if nums[j] == 0 {
				size0++
			} else {
				size1++
			}

			size := size0 + size1

			if size%2 != 0 {
				continue
			}

			threshold := (len(nums) - i) / 2
			if size0 > threshold || size1 > threshold {
				size0 = 0
				size1 = 0
				break
			}

			if size0 == size1 {
				maxSize = max(maxSize, size)
			}
		}
	}

	return maxSize
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
