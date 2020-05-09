package main

func isHappy(n int) bool {
	memo := make(map[int]int)

	for {
		if _, ok := memo[n]; ok {
			return false
		}
		memo[n]++

		var es []int
		for digits := countDigits(n); digits > 0; digits-- {
			base := pow(10, digits-1)
			w := n / base
			es = append(es, w)
			n = n - w*base
		}

		for i, e := range es {
			if i == 0 {
				n = e * e
			} else {
				n = n + e*e
			}
		}

		if n == 1 {
			return true
		}
	}
}

func countDigits(n int) int {
	var cnt int
	for {
		cnt++
		n /= 10
		if n < 1 {
			break
		}
	}
	return cnt
}

func pow(n, x int) int {
	if x == 0 {
		return 1
	}

	acc := n
	for i := 1; i < x; i++ {
		acc *= n
	}
	return acc
}
