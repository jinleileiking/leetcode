package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var found bool

func findTarget(root *TreeNode, k int) bool {
	found = false

	m := make(map[int]bool)

	traverse(root, m, k)
	// spew.Dump(m)
	return found
}

func traverse(n *TreeNode, m map[int]bool, k int) {

	if _, ok := m[n.Val]; ok {
		found = true
		return
	}
	m[k-n.Val] = true
	if n.Left != nil {
		traverse(n.Left, m, k)

	}
	if n.Right != nil {
		traverse(n.Right, m, k)
	}

	return
}
