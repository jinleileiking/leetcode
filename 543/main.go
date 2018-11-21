package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Sum struct {
	cnt   int
	Total int
}

func parse(root *TreeNode) map[int]Sum {
	var ret map[int]Sum
	ret = make(map[int]Sum)
	if root == nil {
		return ret
	}

	level := 0
	parse_inner(root, level, ret)

	return ret
}

func parse_inner(root *TreeNode, level int, ret map[int]Sum) map[int]Sum {

	if root == nil {
		return ret
	}

	t, ok := ret[level]

	if ok {
		t.Total = t.Total + root.Val
		t.cnt++
		// spew.Dump(t)
		ret[level] = t
	} else {
		ret[level] = Sum{
			cnt:   1,
			Total: root.Val,
		}
	}

	parse_inner(root.Left, level+1, ret)
	parse_inner(root.Right, level+1, ret)

	return ret
}

func diameterOfBinaryTree(root *TreeNode) int {

	var left_level, right_level int
	if root.Left != nil {
		left_level = len(parse(root.Left))
	}
	if root.Right != nil {
		right_level = len(parse(root.Right))
	}

	return left_level + right_level + 1
}
