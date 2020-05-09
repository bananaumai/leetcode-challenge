package main

import (
	"fmt"
	"testing"
)

func Test_isCousins(t *testing.T) {
	tts := []struct {
		inputRoot *TreeNode
		inputX    int
		inputY    int
		expected  bool
	}{
		{
			inputRoot: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			inputX:   4,
			inputY:   3,
			expected: false,
		},
		{
			inputRoot: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			inputX:   5,
			inputY:   4,
			expected: true,
		},
		{
			inputRoot: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			inputX:   2,
			inputY:   3,
			expected: false,
		},
	}

	for _, tt := range tts {
		tt := tt

		t.Run(fmt.Sprintf("input root: %v, x: %d, y: %d", tt.inputRoot, tt.inputX, tt.inputY), func(t *testing.T) {
			t.Parallel()

			actual := isCousins(tt.inputRoot, tt.inputX, tt.inputY)

			if tt.expected != actual {
				t.Errorf("expected: %v <=> actual: %v", tt.expected, actual)
			}
		})
	}
}
