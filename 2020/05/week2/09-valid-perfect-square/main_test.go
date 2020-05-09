package template

import (
	"fmt"
	"testing"
)

func Test_isPerfectSquare(t *testing.T) {
	tts := []struct {
		input    int
		expected bool
	}{
		{1, true},
		{2, false},
		{4, true},
		{8, false},
		{9, true},
		{14, false},
		{16, true},
		{17, false},
		{289, true},
		{83521, true},
		{83520, false},
	}

	for _, tt := range tts {
		tt := tt

		t.Run(fmt.Sprintf("input %v", tt.input), func(t *testing.T) {
			t.Parallel()

			actual := isPerfectSquare(tt.input)

			if tt.expected != actual {
				t.Errorf("expected: %v <=> actual: %v", tt.expected, actual)
			}
		})
	}
}
