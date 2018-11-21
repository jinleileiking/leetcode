package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) (p, n *TreeNode) {
	return searchBSTinner(nil, root, val)
}

func searchBSTinner(p, n *TreeNode, val int) (*TreeNode, *TreeNode) {
	if n.Val == val {
		return p, n
	}

	if val > n.Val {
		if n.Right != nil {
			return searchBSTinner(n, n.Right, val)
		} else {
			return nil, nil
		}

	} else {
		if n.Left != nil {
			return searchBSTinner(n, n.Left, val)
		} else {
			return nil, nil
		}
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {

	// 1. search

	if root == nil {
		return nil
	}

	p, found := searchBST(root, key)

	if found == nil {
		// spew.Dump("no found")

		return root
	}

	// if p == nil {
	// 	return nil
	// }

	// spew.Dump(p)
	// spew.Dump(found)

	foundChange := delANode(found)

	// spew.Dump(foundChange)

	if p == nil {
		return foundChange
	}

	if p.Left != nil && p.Left.Val == found.Val {
		p.Left = foundChange
	}
	if p.Right != nil && p.Right.Val == found.Val {
		p.Right = foundChange
	}
	// spew.Dump(found)

	return root
}

func delANode(root *TreeNode) *TreeNode {

	if root.Right == nil && root.Left == nil {
		root = nil
		return root
	}
	// 1. right is nil

	if root.Right == nil {
		return root.Left
	}

	// 2. left is nil

	// spew.Dump(root)
	if root.Left == nil {
		return root.Right
	}

	// 3. left and right all exists

	// 3.1 find the Min
	min := findMin(root.Right)

	minVal := min.Val

	deleteNode(root, min.Val)

	root.Val = minVal

	return root
}

func findMin(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}

	return findMin(root.Left)
}
