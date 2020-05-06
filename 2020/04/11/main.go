package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type path []string

func diameterOfBinaryTree(root *TreeNode) int {
	paths := traverse(root)

	if paths == nil || len(paths) == 0 {
		return 0
	}

	if len(paths) == 1 {
		return len(paths[0])
	}

	var longestPath path
	for _, path := range paths {
		if len(path) > len(longestPath) {
			longestPath = path
		}
	}

	max := len(longestPath)
	for _, path := range paths {
		var commonLen int
		for i := 0; i < len(path); i++ {
			if path[i] == longestPath[i] {
				commonLen++
			} else {
				break
			}
		}
		total := len(path) + len(longestPath) - commonLen * 2
		if total > max {
			max = total
		}
	}

	return max
}

func traverse(node *TreeNode) []path {
	if node == nil {
		return nil
	}

	var paths []path

	for _, direction := range []string{"l", "r"} {
		var target *TreeNode
		if direction == "l" {
			target = node.Left
		} else {
			target = node.Right
		}

		if target == nil {
			continue
		}

		tails := traverse(target)
		if tails != nil {
			for _, path := range tails {
				newPath := append([]string{direction}, path...)
				paths = append(paths, newPath)
			}
		} else {
			paths = append(paths, []string{direction})
		}
	}

	return paths
}
