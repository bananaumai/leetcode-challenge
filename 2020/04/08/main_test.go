package main

import (
	"fmt"
	"strings"
	"testing"
)

func (n *ListNode) String() string {
	next := n

	var vs []string

	for {
		if next == nil {
			break
		}

		vs = append(vs, fmt.Sprintf("%d", next.Val))

		next = next.Next
	}

	return fmt.Sprintf("[%s]", strings.Join(vs, ", "))
}

func Test_middleNode(t *testing.T) {
	tts := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5, 6}, 4},
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
	}

	for _, tt := range tts {
		tt := tt

		t.Run(fmt.Sprintf("input %v", tt.input), func(t *testing.T) {
			t.Parallel()

			var nodes []*ListNode
			for _, n := range tt.input {
				nodes = append(nodes, &ListNode{Val: n})
			}
			t.Logf("nodes: %v", nodes)
			for i := len(nodes) - 1; i > 0; i-- {
				nodes[i-1].Next = nodes[i]
			}
			t.Logf("nodes: %v", nodes)

			head := nodes[0]

			actual := middleNode(head).Val

			if tt.expected != actual {
				t.Errorf("expected: %d <=> actual: %d (%s)", tt.expected, actual, head)
			}
		})
	}
}
