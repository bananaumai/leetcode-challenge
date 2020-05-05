package main

import (
	"fmt"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	tts := []struct {
		input    []string
		expected [][]string
	}{
		{
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{
				{"bat"},
				{"tan", "nat"},
				{"ate", "tea", "eat"},
			},
		},
	}

	for _, tt := range tts {
		tt := tt

		t.Run(fmt.Sprintf("input %v", tt.input), func(t *testing.T) {
			t.Parallel()

			actual := groupAnagrams(tt.input)

			if len(tt.expected) != len(actual) {
				t.Errorf("result length different - expected: %d(%v) <=> actual: %d(%v)",
					len(tt.expected), tt.expected, len(actual), actual)
			} else {
				var cnt int
				for i := 0; i < len(actual); i++ {
					for j := 0; j < len(actual[i]); j++ {
						a := actual[i][j]

						var isFound bool

						for m := 0; m < len(tt.expected); m++ {
							for n := 0; n < len(tt.expected[m]); n++ {
								e := tt.expected[m][n]
								if e == a {
									isFound = true
									cnt++
									break
								}
							}
							if isFound {
								break
							}
						}

						if !isFound {
							t.Errorf("%s is not found", a)
						}
					}
				}

				if cnt != len(tt.input) {
					t.Errorf("anaglams number is not collect - expected: %d <=> actual: %d", len(tt.input), cnt)
				}
			}
		})
	}
}
