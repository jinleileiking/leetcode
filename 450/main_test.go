package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestMain1(t *testing.T) {
	root := &TreeNode{
		Val: 8,
		Right: &TreeNode{
			Val: 55,
			Left: &TreeNode{
				Val: 39,
				Left: &TreeNode{
					Val: 11,
					Right: &TreeNode{
						Val: 23,
					},
				},
			},
		},
	}

	spew.Dump(deleteNode(root, 11))
}

func TestMain2(t *testing.T) {
	root := &TreeNode{
		Val: 8,
	}

	spew.Dump(deleteNode(root, 8))
}

func TestMain3(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	spew.Dump(deleteNode(root, 3))
}

func TestMain4(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}

	spew.Dump(deleteNode(root, 1))
}

func TestMain5(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}

	spew.Dump(deleteNode(root, 3))
}
