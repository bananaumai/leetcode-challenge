package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	type query struct {
		node   *TreeNode
		parent *TreeNode
		depth  int
	}

	var found []*query
	for qs := []*query{{root, nil, 0}}; len(qs) > 0; {
		q := qs[0]
		qs = append(make([]*query, 0), qs[1:]...)

		if q.node == nil {
			continue
		}

		if q.node.Val == x || q.node.Val == y {
			found = append(found, q)

			if len(found) == 2 {
				q1 := found[0]
				q2 := found[1]
				return q1.depth == q2.depth && q1.parent.Val != q2.parent.Val
			}
		}

		qs = append(qs, &query{q.node.Left, q.node, q.depth + 1})
		qs = append(qs, &query{q.node.Right, q.node, q.depth + 1})
	}

	return false
}
