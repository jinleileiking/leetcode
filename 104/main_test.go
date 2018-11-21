package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestMain(t *testing.T) {
	root := &TreeNode{Val: 2}
	left := &TreeNode{Val: 1}
	right := &TreeNode{Val: 3}

	root.Left = left
	root.Right = right
	spew.Dump(isValidBST(root))

}

func TestMain1(t *testing.T) {
	root := &TreeNode{Val: 2}
	left := &TreeNode{Val: 3}
	right := &TreeNode{Val: 1}

	root.Left = left
	root.Right = right
	spew.Dump(isValidBST(root))

}

func TestMain2(t *testing.T) {
	root := &TreeNode{Val: 1}
	left := &TreeNode{Val: 1}

	root.Left = left
	spew.Dump(isValidBST(root))

}

func TestMain3(t *testing.T) {
	root := &TreeNode{Val: 1}
	right := &TreeNode{Val: 1}

	root.Right = right
	spew.Dump(isValidBST(root))

}
