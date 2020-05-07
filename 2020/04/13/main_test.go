package main

import (
	"fmt"
	"testing"
)

func Test_findMaxLength(t *testing.T) {
	tts := []struct {
		input    []int
		expected int
	}{
		{[]int{0, 1}, 2},
		{[]int{0, 1, 0}, 2},
		{[]int{0, 1, 1, 0, 0}, 4},
		{[]int{0, 0, 0, 1, 1, 1, 0}, 6},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
		{[]int{0, 1, 1, 1, 0, 1, 0, 1, 0}, 6},
		{[]int{0, 1, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0}, 12},
	}

	for _, tt := range tts {
		tt := tt

		t.Run(fmt.Sprintf("input %v", tt.input), func(t *testing.T) {
			t.Parallel()

			actual := findMaxLength(tt.input)

			if tt.expected != actual {
				t.Errorf("expected: %d <=> actual: %d", tt.expected, actual)
			}
		})
	}
}
