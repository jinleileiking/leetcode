package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(node *TreeNode) {
	if node == nil {
		return
	}

	node_new := &TreeNode{
		Val: node.Val,
	}

	if ret == nil {
		ret = node_new
		ret.Right = nil
		rightnode = ret
	} else {
		rightnode.Right = node_new
		rightnode = rightnode.Right
	}

	// spew.Dump(ret)
	traverse(node.Left)
	traverse(node.Right)

}

var ret *TreeNode
var rightnode *TreeNode

func flatten(root *TreeNode) {
	ret = nil
	rightnode = nil

	if root == nil {
		return
	}
	traverse(root)

	// spew.Dump(ret)
	// root = ret
	// spew.Dump(root)

	root.Left = nil
	root.Right = ret.Right

}
