package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestMain(t *testing.T) {
	root := &TreeNode{Val: 2}
	left := &TreeNode{Val: 3}
	right := &TreeNode{Val: 1}

	root.Left = left
	root.Right = right

	flatten(root)

	spew.Dump(root)
}

func TestMain1(t *testing.T) {
	root := &TreeNode{Val: 0}

	flatten(root)

	spew.Dump(root)
}
