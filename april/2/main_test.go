package main

import (
	"fmt"
	"testing"
)

func Test_isHappy(t *testing.T) {
	tts := []struct {
		input int
		expected bool
	} {
		//{19, true},
		{1221, true},
	}

	for _, tt := range tts {
		tt := tt

		t.Run(fmt.Sprintf("input: %d", tt.input), func(t *testing.T) {
			t.Parallel()

			actual := isHappy(tt.input)

			if tt.expected != actual {
				t.Errorf("expected: %v <=> actual: %v", tt.expected, actual)
			}
		})
	}
}
