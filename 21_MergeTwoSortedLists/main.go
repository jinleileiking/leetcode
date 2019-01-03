package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		temp := l1
		l1 = l2
		l2 = temp
	}

	start := l1

	// l2 is bigger
	for l1 != nil && l2 != nil {
		// spew.Dump(".....................", start)
		// spew.Dump(l1)

		// 3 4 7
		// 5 6 8
		if l1.Val <= l2.Val {

			for l1.Next != nil && l1.Next.Val <= l2.Val {
				l1 = l1.Next
			}
			// l1 -> 4

			l2Start := l2
			// var gotNext bool
			for l2.Next != nil && l1.Next != nil &&
				l2.Val >= l1.Val && l2.Next.Val <= l1.Next.Val {
				l2 = l2.Next
			}

			l1Next := l1.Next

			l1.Next = l2Start

			l2Next := l2.Next

			l2.Next = l1Next

			l2 = l2Next

		} else {
			l1.Next = l1
		}
	}

	if l2 != nil && l1 == nil {
		l1.Next = l2
	}
	return start
}

func mergeTest(l1 *ListNode, l2 *ListNode) *ListNode {

	// l1, 1,  3
	// l2  2,

	temp := l1.Next
	l1.Next = l2
	l2.Next = temp

	return l1

}
