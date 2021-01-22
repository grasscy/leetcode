package sort

import (
	"testing"
)

func Test_insertionSortList(t *testing.T) {
	n1 := ListNode{
		Val:  3,
		Next: nil,
	}
	n2 := ListNode{
		Val:  1,
		Next: &n1,
	}
	n3 := ListNode{
		Val:  2,
		Next: &n2,
	}
	n4 := ListNode{
		Val:  4,
		Next: &n3,
	}

	insertionSortList(&n4)

	n5 := ListNode{
		Val:  0,
		Next: nil,
	}
	n6 := ListNode{
		Val:  4,
		Next: &n5,
	}
	n7 := ListNode{
		Val:  3,
		Next: &n6,
	}
	n8 := ListNode{
		Val:  5,
		Next: &n7,
	}
	n9 := ListNode{
		Val:  -1,
		Next: &n8,
	}

	insertionSortList(&n9)
}
