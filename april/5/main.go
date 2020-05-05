package main

func maxProfit(prices []int) int {
	return newSolver(prices).solve(0)
}

const dpInit = -1

type solver struct {
	prices []int
	dp []int
}

func newSolver(prices []int) solver {
	dpSize := len(prices)
	dp := make([]int, dpSize)
	for i := 0; i < dpSize; i++ {
		dp[i] = dpInit
	}
	return solver{prices, dp}
}

func (s solver) solve(start int) int {
	answer := s.dp[start]
	if answer != dpInit {
		return answer
	}

	max := 0
	for i := start; i < len(s.prices); i++ {
		for j := i + 1; j < len(s.prices); j++ {
			profit := s.prices[j] - s.prices[i]

			if j + 1 < len(s.prices) {
				rest := s.solve(j+1)
				if profit >= 0 {
					profit += rest
				} else {
					profit = rest
				}
			}

			if profit > max {
				max = profit
			}
		}
	}
	s.dp[start] = max
	return max
}
