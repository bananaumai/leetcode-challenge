package main

import (
	"fmt"
	"testing"
)

func Test_canConstruct(t *testing.T) {
	tts := []struct{
		r string
		m string
		expected bool
	} {
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
		{"bg", "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj", true},
	}

	for _, tt := range tts {
		tt := tt
		t.Run(fmt.Sprintf("ransomNotes: %s, magazine: %s", tt.r, tt.m), func(t *testing.T) {
			actual := canConstruct(tt.r, tt.m)
			if actual != tt.expected {
				t.Errorf("expected: %v <=> actual: %v", tt.expected, actual)
			}
		})
	}
}
