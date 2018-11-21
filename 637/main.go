package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var level int

type Sum struct {
	cnt   int
	Total int
}

var ret map[int]Sum

func averageOfLevels(root *TreeNode) []float64 {

	ret = make(map[int]Sum)

	average(root, 0)

	// spew.Dump(ret)

	ret_float := make([]float64, len(ret))

	for k, v := range ret {
		// ret_float = append(ret_float, float64(v.Total)/float64(v.cnt))
		ret_float[k] = float64(v.Total) / float64(v.cnt)

		// spew.Dump(ret_float)
	}
	return ret_float
}

func average(node *TreeNode, level int) {
	if node == nil {
		return
	}

	t, ok := ret[level]

	// spew.Dump(level)
	// spew.Dump(ok)
	// spew.Dump(node.Val)
	if ok {
		t.Total = t.Total + node.Val
		t.cnt++
		// spew.Dump(t)
		ret[level] = t
	} else {
		ret[level] = Sum{
			cnt:   1,
			Total: node.Val,
		}
	}

	average(node.Left, level+1)
	average(node.Right, level+1)
}
