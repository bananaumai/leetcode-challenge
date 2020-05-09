package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	var nodes []*ListNode

	next := head
	for {
		if next == nil {
			break
		}

		nodes = append(nodes, next)

		next = next.Next
	}

	if nodes == nil {
		return nil
	}

	size := len(nodes)

	if size == 0 {
		return nil
	}

	if size == 1 {
		return nodes[0]
	}

	return nodes[size/2]
}
