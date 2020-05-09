package main

import (
	"fmt"
	"testing"
)

func Test_findComplement(t *testing.T) {
	tts := []struct {
		input    int
		expected int
	}{
		{5, 2},
		{1, 0},
	}

	for _, tt := range tts {
		tt := tt
		t.Run(fmt.Sprintf("input %d", tt.input), func(t *testing.T) {
			t.Parallel()

			actual := findComplement(tt.input)

			if tt.expected != actual {
				t.Errorf("expected: %d <=> actual: %d", tt.expected, actual)
			}
		})
	}
}
