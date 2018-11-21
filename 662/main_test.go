package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestMain(t *testing.T) {
	root := &TreeNode{Val: 3}
	left := &TreeNode{Val: 4}
	right := &TreeNode{Val: 5}

	root.Left = left
	root.Right = right

	// spew.Dump(widthOfBinaryTree(root))
	spew.Dump(widthOfBinaryTree(root))
}
