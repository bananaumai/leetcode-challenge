package main

import "sort"

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)
	for {
		cnt := len(stones)

		if cnt == 0 {
			return 0
		}

		if cnt == 1 {
			return stones[0]
		}

		x := stones[cnt-2]
		y := stones[cnt-1]

		stones = append(make([]int, 0), stones[:cnt-2]...)

		if x != y {
			n := y - x
			stones = append(stones, n)
			sort.Ints(stones)
		}
	}
}
