package main

import (
	"strings"
)

func backspaceCompare(S string, T string) bool {
	return parse(S) == parse(T)
}

func parse(s string) string {
	var pos int
	var rs []rune
	for _, r := range s {
		// # is 35
		if r != 35 {
			rs = append(rs, r)
			pos++
		} else {
			if pos != 0 {
				rs = append(make([]rune, 0), rs[:pos-1]...)
				pos--
			}
		}
	}

	b := &strings.Builder{}
	for _, r := range rs {
		b.WriteRune(r)
	}
	return b.String()
}
