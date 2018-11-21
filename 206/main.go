package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var prev, cur, ret *ListNode
	cur = head
	for {
		// spew.Dump(ret, "-----------------")
		// spew.Dump("xxxxx")
		if prev == nil {
			prev = head
			cur = head.Next
			head.Next = nil
			continue
		}

		// spew.Dump(cur, prev, "-----------------")
		temp := cur.Next
		cur.Next = prev
		ret = cur
		cur = temp
		prev = ret

		if prev == head {
			// spew.Dump("aaaaaaaaaaa")
			head.Next = nil
		}

		if cur == nil {
			// spew.Dump(cur, ret, prev)
			// prev.Next = nil
			break
		}

	}

	return ret
}
