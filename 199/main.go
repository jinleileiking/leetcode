package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(node *TreeNode, ret *[]int, level int) {
	if node == nil {
		return
	}

	if len(*ret) == level {
		*ret = append(*ret, node.Val)
	}

	level++
	traverse(node.Right, ret, level)
	traverse(node.Left, ret, level)

}

func rightSideView(root *TreeNode) []int {
	ret := &[]int{}
	level := 0

	traverse(root, ret, level)

	return *ret
}
