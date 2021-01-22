package sort

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	res := head

	for cur := head; cur != nil; cur = cur.Next {
		c := res
		for ; c.Next != nil && c != cur && c.Next.Val < cur.Val; c = c.Next {
		}
		temp := c.Next
		c.Next = cur
		cur.Next = temp
	}
	return head
}
