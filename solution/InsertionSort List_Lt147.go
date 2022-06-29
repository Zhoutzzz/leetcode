package main

func insertionSortList(head *ListNode) *ListNode {
	pre, cur, next := head, head, head.Next
	for cur != nil {
		if cur.Val < pre.Val {
			next = cur.Next
			var tmp, p = head, &ListNode{}
			for tmp.Val < cur.Val {
				p = tmp
				tmp = tmp.Next
			}
			p.Next = cur
			cur.Next = tmp
		}
		pre = cur
		cur = next
	}

	return head
}
