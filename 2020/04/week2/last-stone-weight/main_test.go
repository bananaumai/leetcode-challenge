package main

import (
	"fmt"
	"testing"
)

func Test_lastStoneWeight(t *testing.T) {
	tts := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
		{[]int{1, 3}, 2},
	}

	for _, tt := range tts {
		tt := tt

		t.Run(fmt.Sprintf("input %v", tt.input), func(t *testing.T) {
			t.Parallel()

			actual := lastStoneWeight(tt.input)

			if tt.expected != actual {
				t.Errorf("expected: %d <=> actual: %d", tt.expected, actual)
			}
		})
	}
}
