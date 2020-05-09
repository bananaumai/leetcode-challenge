package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	c, m := diameter(root)
	return max(c-1, m)
}

func diameter(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}

	lCount, lMax := diameter(node.Left)
	rCount, rMax := diameter(node.Right)
	return max(lCount, rCount) + 1, max(lCount+rCount, max(lMax, rMax))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
