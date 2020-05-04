package main

import (
	"fmt"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	tts := []struct {
		input []int
		expected int
	} {
		{[]int{-2,1,-3,4,-1,2,1,-5,4}, 6},
	}

	for _, tt := range tts {
		tt := tt
		t.Run(fmt.Sprintf("input: %v", tt.input), func(t *testing.T) {
			t.Parallel()

			actual := maxSubArray(tt.input)

			if tt.expected != actual {
				t.Errorf("expected: %d <=> actual: %d", tt.expected, actual)
			}
		})
	}
}
