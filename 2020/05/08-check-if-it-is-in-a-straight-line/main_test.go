package main

import (
	"fmt"
	"testing"
)

func Test_checkStraightLine(t *testing.T) {
	tts := []struct {
		input    [][]int
		expected bool
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {6, 7}}, true},
		{[][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}, false},
	}

	for _, tt := range tts {
		tt := tt

		t.Run(fmt.Sprintf("input %v", tt.input), func(t *testing.T) {
			t.Parallel()

			actual := checkStraightLine(tt.input)

			if tt.expected != actual {
				t.Errorf("expected: %v <=> actual: %v", tt.expected, actual)
			}
		})
	}
}
