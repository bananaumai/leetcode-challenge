package main

func moveZeroes(nums []int)  {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			chunk := 1
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					break
				}
				chunk++
			}
			swap(nums, i, chunk)
		}
	}
}

func swap(arr []int, pos, size int) {
	for i := pos; i < pos + size; i++ {
		target := i+size
		if target >= len(arr) {
			break
		}
		tmp := arr[target]
		arr[target] = arr[i]
		arr[i] = tmp
	}
}
