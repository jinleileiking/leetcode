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

	leftleft := &TreeNode{Val: 1}
	leftright := &TreeNode{Val: 2}

	left.Left = leftleft
	left.Right = leftright

	root1 := &TreeNode{Val: 4}
	left1 := &TreeNode{Val: 1}
	right1 := &TreeNode{Val: 2}

	root1.Left = left1
	root1.Right = right1

	spew.Dump(isSubtree(root, root1))
}

func TestMain1(t *testing.T) {
	root := &TreeNode{Val: 3}
	left := &TreeNode{Val: 4}
	right := &TreeNode{Val: 5}

	root.Left = left
	root.Right = right

	leftleft := &TreeNode{Val: 1}
	leftright := &TreeNode{Val: 2}

	left.Left = leftleft
	left.Right = leftright

	root1 := &TreeNode{Val: 4}
	left1 := &TreeNode{Val: 1}
	right1 := &TreeNode{Val: 2}

	root1.Left = left1
	root1.Right = right1

	spew.Dump(isSubtree(root, root1))
}

func TestContain(t *testing.T) {

	s := []int{1, 2, 3, 4, 5}
	d := []int{2, 3, 4}
	spew.Dump(contain(s, d))
}

func TestContain1(t *testing.T) {

	s := []int{1, 2, 3, 4, 5}
	d := []int{2, 3, 5}
	spew.Dump(contain(s, d))
}

func TestContain2(t *testing.T) {

	s := []int{1, 2, 3, 4, 5}
	d := []int{3, 4, 5}
	spew.Dump(contain(s, d))
}

func TestContain3(t *testing.T) {

	s := []int{1, 2, 3, 4, 5}
	d := []int{1, 2, 3}
	spew.Dump(contain(s, d))
}

func TestContain4(t *testing.T) {

	s := []int{1, 2, 3}
	d := []int{1, 2, 3}
	spew.Dump(contain(s, d))
}
