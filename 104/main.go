package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMaxDepth(root *TreeNode, currentMaxDepth *int, currentDepth int) {
	currentDepth++
	if root.Left == nil && root.Right == nil {
		if currentDepth > *currentMaxDepth {
			*currentMaxDepth = currentDepth
		}
		return
	}

	for _, v := range []*TreeNode{root.Left, root.Right} {
		if v != nil {
			findMaxDepth(v, currentMaxDepth, currentDepth)
		}
	}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var currentMaxDepth int
	findMaxDepth(root, &currentMaxDepth, 0)
	return currentMaxDepth
}
