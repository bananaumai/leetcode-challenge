package main

import (
	"fmt"
	"testing"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	tts := []struct {
		input *TreeNode
		expected int
	} {
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{Val: 3},
			},
			expected: 3,
		},
		{
			input: &TreeNode{},
			expected: 0,
		},
		{
			input: nil,
			expected: 0,
		},
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{Val: 2},
			},
			expected: 1,
		},
		{
			input: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{Val: 1},
				},
			},
			expected: 2,
		},
		{
			input: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{Val: 1},
						Right: &TreeNode{Val: 2},
					},
				},
			},
			expected: 3,
		},
		{
			input: &TreeNode{
				Val: 4,
				Left: &TreeNode{Val: -7},
				Right: &TreeNode{
					Val: -3,
					Left: &TreeNode{
						Val: -9,
						Left: &TreeNode{
							Val: 9,
							Left: &TreeNode{
								Val: 6,
								Left: &TreeNode{
									Val: 0,
									Left: &TreeNode{
										Val: -1,
									},
								},
								Right: &TreeNode{
									Val: 6,
									Left: &TreeNode{
										Val: 4,
									},
								},
							},
						},
						Right: &TreeNode{
							Val: -7,
							Left: &TreeNode{
								Val: -6,
								Left: &TreeNode{
									Val: 5,
								},
							},
							Right: &TreeNode{
								Val: -6,
								Left: &TreeNode{
									Val: 9,
									Left: &TreeNode{
										Val: -2,
									},
								},
							},
						},
					},
					Right: &TreeNode{
						Val: -3,
						Left: &TreeNode{
							Val: -4,
						},
					},
				},
			},
			expected: 8,
		},
	}

	for _, tt := range tts {
		tt := tt

		t.Run(fmt.Sprintf("input %v", tt.input), func(t *testing.T) {
			t.Parallel()

			actual := diameterOfBinaryTree(tt.input)

			if tt.expected != actual {
				t.Errorf("expected: %d <=> actual: %d", tt.expected, actual)
			}
		})
	}
}
