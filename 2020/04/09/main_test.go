package main

import (
	"fmt"
	"testing"
)

func Test_backspaceCompare(t *testing.T) {
	tts := []struct {
		inputS string
		inputT string
		expected bool
	} {
		{"ab#c", "ad#c", true},
		{"ab##", "c#d#", true},
		{"a##c", "#a#c", true},
		{"a#c", "b", false},
	}

	for _, tt := range tts {
		tt := tt

		t.Run(fmt.Sprintf("input s:%s t:%s", tt.inputS, tt.inputT), func(t *testing.T) {
			t.Parallel()

			actual := backspaceCompare(tt.inputS, tt.inputT)

			if tt.expected != actual {
				t.Errorf("expected: %v <=> actual: %v", tt.expected, actual)
			}
		})
	}
}
