package main

import (
	"fmt"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	tts := []struct {
		input []int
		expected int
	} {
		{[]int{2,2,1}, 1},
		{[]int{4,1,2,1,2}, 4},
	}

	for _, tt := range tts {
		tt := tt
		t.Run(fmt.Sprintf("input: %v", tt.input), func(t *testing.T) {
			t.Parallel()

			actual := singleNumber(tt.input)

			if tt.expected != actual {
				t.Errorf("expected: %d <=> actual: %d", tt.expected, actual)
			}
		})
	}
}
