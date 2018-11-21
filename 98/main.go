package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	if root == nil {
		return true
	}

	left := root.Left
	right := root.Right

	if left == nil && right == nil {
		return true
	}

	if left != nil && left.Val >= root.Val {
		return false
	}

	if right != nil && right.Val <= root.Val {
		return false
	}

	if false == isValidBST(left) {
		return false
	}

	if false == isValidBST(right) {
		return false
	}

	return true

}
