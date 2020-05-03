package main

func canConstruct(ransomNote string, magazine string) bool {
	ransomMap := make(map[int32]int)
	for _, c := range ransomNote {
		ransomMap[c]++
	}

	magazineMap := make(map[int32]int)
	for _, c := range magazine {
		magazineMap[c]++
	}

	for c, rCnt := range ransomMap {
		mCnt, ok := magazineMap[c]
		if !ok {
			return false
		}
		if mCnt < rCnt {
			return false
		}
	}

	return true
}
