package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestTest(t *testing.T) {

	l1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
		},
	}

	l2 := ListNode{
		Val: 2,
	}

	spew.Dump(mergeTest(&l1, &l2))
	// assert.Equal(t, 3, find("abcabcbb", 3), "they should be equal")
}

func TestMain(t *testing.T) {

	l1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
			},
		},
	}

	l2 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 6,
			},
		},
	}

	spew.Dump(mergeTwoLists(&l1, &l2))
	// assert.Equal(t, 3, find("abcabcbb", 3), "they should be equal")
}

func TestMain1(t *testing.T) {

	l1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 5,
		},
	}

	l2 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 6,
			},
		},
	}

	spew.Dump(mergeTwoLists(&l1, &l2))
	// assert.Equal(t, 3, find("abcabcbb", 3), "they should be equal")
}

func TestMain2(t *testing.T) {

	l1 := ListNode{
		Val: -8,
		Next: &ListNode{
			Val: -1,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l2 := ListNode{
		Val: -8,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 8,
								Next: &ListNode{
									Val: 9,
								},
							},
						},
					},
				},
			},
		},
	}

	spew.Dump(mergeTwoLists(&l1, &l2))
	// assert.Equal(t, 3, find("abcabcbb", 3), "they should be equal")
}

func TestMain3(t *testing.T) {

	l1 := ListNode{
		Val: 5,
	}

	l2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	spew.Dump(mergeTwoLists(&l1, &l2))
	// assert.Equal(t, 3, find("abcabcbb", 3), "they should be equal")
}

func TestMain4(t *testing.T) {

	l1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l2 := ListNode{
		Val: 5,
	}

	spew.Dump(mergeTwoLists(&l1, &l2))
	// assert.Equal(t, 3, find("abcabcbb", 3), "they should be equal")
}
