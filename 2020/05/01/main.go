package main

import "fmt"

var (
	n          = 3
	badVersion = 2
	apiCallCnt int
)

func main() {
	answer := firstBadVersion(n)
	fmt.Printf("answer: %d, api call cnt: %d\n", answer, apiCallCnt)
}

func firstBadVersion(n int) int {
	return solve(0, n, 0)
}

func solve(start int, end int, answer int) int {
	if start == end {
		if isBadVersion(start) {
			return start
		}
		return answer
	}

	if start+1 == end {
		if isBadVersion(start) {
			return start
		}
		if isBadVersion(end) {
			return end
		}
		return answer
	}

	half := (start + end) / 2

	if isBadVersion(half) {
		return solve(start, half-1, half)
	}

	return solve(half+1, end, answer)
}

func isBadVersion(version int) bool {
	apiCallCnt++
	return version >= badVersion
}
