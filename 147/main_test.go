package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestMain1(t *testing.T) {
	root := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 55,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 11,
					},
				},
			},
		},
	}

	spew.Dump(insertionSortList(root))
}

func TestMain2(t *testing.T) {
	root := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}

	spew.Dump(insertionSortList(root))
}

func TestMain3(t *testing.T) {
	root := &ListNode{
		Val: -1,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 0,
					},
				},
			},
		},
	}

	spew.Dump(insertionSortList(root))
}

func TestMain4(t *testing.T) {
	root := &ListNode{
		Val: -1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 0,
					},
				},
			},
		},
	}

	spew.Dump(insertionSortList(root))
}
