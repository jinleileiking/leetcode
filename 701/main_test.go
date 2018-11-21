package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestMain0(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
		},
	}

	spew.Dump(insertIntoBST(root, 5))

}

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

	spew.Dump(insertIntoBST(root, 17))

}

func TestMain2(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Right: &TreeNode{
			Val: 14,
			Left: &TreeNode{
				Val: 10,
			},
			Right: &TreeNode{
				Val: 77,
				Right: &TreeNode{
					Val: 95,
				},
			},
		},
	}

	spew.Dump(insertIntoBST(root, 4))

}
