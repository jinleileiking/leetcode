package main

import (
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(t *TreeNode, ret *[]int) {
	if t == nil {
		return
	}
	*ret = append(*ret, t.Val)
	// spew.Dump(ret)
	traverse(t.Left, ret)
	traverse(t.Right, ret)
}

func toSlice(t *TreeNode) []int {

	ret := &[]int{}
	traverse(t, ret)
	// spew.Dump(ret)

	return *ret
}

func contain(s, d []int) bool {

	small := len(d)

	for i := 0; i < len(s)-small+1; i++ {
		// spew.Dump(i)
		if reflect.DeepEqual(s[i:i+len(d)], d) {
			return true
		}
	}

	return false

}

func equal(s, d []int) bool {

	if len(s) != len(d) {
		return false
	}

	return reflect.DeepEqual(s, d)

}
func traverse1(t *TreeNode, d []int) bool {
	if t == nil {
		return false
	}
	n_slice := toSlice(t)

	if equal(n_slice, d) {
		return true
	}
	// spew.Dump(ret)
	if traverse1(t.Left, d) {
		return true
	}
	if traverse1(t.Right, d) {
		return true
	}
	return false
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	t_slice := toSlice(t)

	return traverse1(s, t_slice)
	// spew.Dump(s_slice)
	// spew.Dump(t_slice)
}
