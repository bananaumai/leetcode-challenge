package main

import "fmt"

func main() {
	J := "z"
	S := "ZZ"
	answer := numJewelsInStones(J, S)
	fmt.Printf("answer: %d", answer)
}

func numJewelsInStones(J string, S string) int {
	js := make(map[int32]bool)
	for _, c := range J {
		js[c] = true
	}

	var cnt int
	for _, c := range S {
		if ok, _ := js[c]; ok {
			cnt++
		}
	}

	return cnt
}
