package main

import "github.com/davecgh/go-spew/spew"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	// spew.Dump(root.Val)
	spew.Dump(1)
	if root == nil {
		return nil
	}

	right := invertTree(root.Right)
	left := invertTree(root.Left)
	// spew.Dump(root.Val)

	root.Right = left
	root.Left = right

	return root

}
