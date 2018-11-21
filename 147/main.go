package main

import "github.com/davecgh/go-spew/spew"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swap(n1, n2 *ListNode) {

	tmp := n1.Val

	n1.Val = n2.Val

	n2.Val = tmp

}

func insertionSortList(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	var tmp int

	for end := head.Next; end != nil; end = end.Next {
		// spew.Dump(end)
		needSwap := false
		judgeVal := end.Val
		for node := head; node != end; node = node.Next {
			if needSwap {
				// spew.Dump(head)
				// spew.Dump(tmp)
				t := node.Val
				node.Val = tmp
				tmp = t
				continue
			}
			if judgeVal < node.Val {
				tmp = node.Val
				node.Val = judgeVal
				needSwap = true
			}
		}

		if needSwap {
			end.Val = tmp
		}
		spew.Dump(head)
	}

	return head
}
