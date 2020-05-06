package main

func firstUniqChar(s string) int {
	cMap := make(map[int32]int)

	for _, c := range s {
		cMap[c]++
	}

	for i, c := range s {
		cnt, _ := cMap[c]
		if cnt == 1 {
			return i
		}
	}

	return -1
}
