package main

import (
	"fmt"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	tts := []struct {
		input []int
		expected []int
	}{
	 	//{[]int{0,1,0,3,12}, []int{1,3,12,0,0}},
	 	{[]int{4,2,4,0,0,3,0,5,1,0}, []int{4,2,4,3,5,1,0,0,0,0}},
	}

	for _, tt := range tts {
		tt := tt
		t.Run(fmt.Sprintf("input %v", tt.input), func(t *testing.T) {
			t.Parallel()

			moveZeroes(tt.input)

			for i := 0; i < len(tt.input); i++ {
				if tt.expected[i] != tt.input[i] {
					t.Errorf("unexpected result at %d - expected: %d <=> actual: %d", i, tt.expected[i], tt.input[i])
				}
			}
		})
	}
}
