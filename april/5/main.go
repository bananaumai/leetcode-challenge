package main

func maxProfit(prices []int) int {
	return newSolver(prices).solve(0)
}

const dpInit = -1

type solver struct {
	prices []int
	dp [][]int
}

func newSolver(prices []int) solver {
	dpSize := len(prices)
	dp := make([][]int, dpSize)
	for i := 0; i < dpSize; i++ {
		dp[i] = make([]int, dpSize)
		for j := 0; j < dpSize; j++ {
			dp[i][j] = dpInit
		}
	}
	return solver{prices, dp}
}

func (s solver) solve(start int) int {
	max := 0
	for i := start; i < len(s.prices); i++ {
		for j := i + 1; j < len(s.prices); j++ {
			if s.dp[i][j] == dpInit {
				profit := s.prices[j] - s.prices[i]
				if profit >= 0 {
					profit += s.solve(j+1)
				} else {
					profit = s.solve(j+1)
				}

				if profit >= 0 {
					s.dp[i][j] = profit
				} else {
					s.dp[i][j] = 0
				}
			}

			if s.dp[i][j] > max {
				max = s.dp[i][j]
			}
		}
	}
	return max
}
