package main

func findComplement(num int) int {
	d := digits(num)
	p := pow(2, d + 1)
	return num ^ (p - 1)
}

func digits(num int) int {
	var cnt int
	for num > 1 {
		cnt += 1
		num /= 2
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
