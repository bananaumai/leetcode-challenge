package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)
	for _, str := range strs {
		key := normalize(str)
		strMap[key] = append(strMap[key], str)
	}

	var result [][]string
	for _, ss := range strMap {
		result = append(result, ss)
	}

	return result
}

func normalize(str string) string {
	var cs []string

	for _, c := range str {
		cs = append(cs, fmt.Sprintf("%d", c))
	}

	sort.Strings(cs)

	return strings.Join(cs, "_")
}
