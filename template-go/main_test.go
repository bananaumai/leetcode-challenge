package main

import (
	"fmt"
	"testing"
)

func Test_foo(t *testing.T) {
	tts := []struct {
		input    interface{}
		expected interface{}
	}{
		{},
	}

	for _, tt := range tts {
		tt := tt

		t.Run(fmt.Sprintf("input %v", tt.input), func(t *testing.T) {
			t.Parallel()

			actual := foo(t)

			if tt.expected != actual {
				t.Errorf("expected: %v <=> actual: %v", tt.expected, actual)
			}
		})
	}
}
