package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	paths := traverse(root)

	if paths == nil || len(paths) == 0 {
		return 0
	}

	if len(paths) == 1 {
		return len(paths[0])
	}

	pairs := combination(paths)

	var max int
	for _, pair := range pairs {
		path1 := pair[0]
		path1Len := len(path1)
		if path1Len > max {
			max = path1Len
		}

		path2 := pair[1]
		path2Len := len(path2)
		if path2Len > max {
			max = path2Len
		}

		var commonLen int
		for i := 0; i < len(path1); i++ {
			if path1[i] == path2[i] {
				commonLen++
			} else {
				break
			}
		}
		total := path1Len + path2Len - commonLen * 2
		if total > max {
			max = total
		}
	}

	return max
}

type path []string

func combination(paths []path) [][]path {
	var pairs [][]path

	for i := 0; i < len(paths); i++ {
		for j := i + 1; j < len(paths); j++ {
			pair := []path{paths[i], paths[j]}
			pairs = append(pairs, pair)
		}
	}

	return pairs
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
