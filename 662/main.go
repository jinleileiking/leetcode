package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var width map[int]int

func traverse(t *TreeNode, level int) {
	if t == nil {
		return
	}

	if t.Left != nil && t.Right == nil {
		width[level+1]++
	}

	if t.Right != nil && t.Left == nil {
		width[level+1]++
	}

	width[level]++
	level++
	traverse(t.Left, level)
	traverse(t.Right, level)
}

func widthOfBinaryTree(root *TreeNode) int {

	width = make(map[int]int, 0)

	level := 0
	traverse(root, level)

	// spew.Dump(width)

	var max int
	for _, v := range width {
		if v > max {
			max = v
		}
	}

	return max
}
