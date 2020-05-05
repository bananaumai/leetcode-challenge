package main

import (
	"fmt"
	"testing"
)

func Test_firstUniqChar(t *testing.T) {
	tts := []struct {
		input string
		expected int
	} {
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aa", -1},
	}

	for _, tt := range tts {
		tt := tt

		t.Run(fmt.Sprintf("input %s", tt.input), func(t *testing.T) {
			t.Parallel()

			actual := firstUniqChar(tt.input)

			if tt.expected != actual {
				t.Errorf("expected: %d <=> actual: %d", tt.expected, actual)
			}
		})
	}
}
