package main

import "github.com/davecgh/go-spew/spew"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {

	spew.Dump(root)

	insert := TreeNode{
		Val: val,
	}

	if root != nil && root.Left == nil && root.Right == nil {
		if root.Val > val {
			root.Left = &insert
		} else {
			root.Right = &insert
		}
		return root
	}

	if root != nil && root.Val > val {
		if root.Left == nil {
			root.Left = &insert
			return root
		}
		insertIntoBST(root.Left, val)
	} else {
		if root.Right == nil {
			root.Right = &insert
			return root
		}
		insertIntoBST(root.Right, val)
	}
	// if root.Left != nil && root.Left.Val < val {
	// 	if root.Right == nil {
	// 		root.Right = &insert
	// 		return root
	// 	}
	// 	insertIntoBST(root.Right, val)
	// } else if root.Right != nil && root.Right.Val > val {
	// 	if root.Left == nil {
	// 		root.Left = &insert
	// 		return root
	// 	}
	// 	insertIntoBST(root.Left, val)
	// }

	return root
}
