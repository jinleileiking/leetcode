package main

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getTier(list *ListNode) int {
	var tier int

	if list.Next == nil {
		return 0
	}

	for list.Next != nil {
		tier++
		list = list.Next
	}

	return tier
}

func getLen(num int) int {
	for i := 0; ; i++ {
		if num/int(math.Pow(float64(10), float64(i))) == 0 {
			return i
		}
	}
}

func genLink(num int) *ListNode {

	var retnode, prenode *ListNode

	// start := false

	// spew.Dump(getLen(num))

	if num == 0 {
		return &ListNode{Val: 0}
	}
	for i := 0; i < getLen(num); i++ {
		// spew.Dump(i)
		node := &ListNode{}

		if prenode == nil {
			retnode = node
		}

		node.Val = num / int(math.Pow(float64(10), float64(i))) % 10

		if prenode != nil {
			prenode.Next = node
		}
		prenode = node
	}

	return retnode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	tier1 := getTier(l1)
	tier2 := getTier(l2)

	var max int
	if tier1 > tier2 {
		max = tier1
	} else {
		max = tier2
	}

	var num1, num2 int
	for level := 0; level < max+1; level++ {
		if l1 != nil {
			num1 = num1 + int(math.Pow(float64(10), float64(level)))*l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			num2 = num2 + int(math.Pow(float64(10), float64(level)))*l2.Val
			l2 = l2.Next
		}
	}

	total := num1 + num2
	// fmt.Println(total)
	return genLink(total)
}
