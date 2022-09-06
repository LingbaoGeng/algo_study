package link

import "algo_study/object"

func ReverseKGroup(head *object.ListNode, k int) *object.ListNode {
	if head == nil {
		return nil
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverse(a, b)
	a.Next = ReverseKGroup(b, k)
	return newHead
}

func reverse(a, b *object.ListNode) *object.ListNode {
	var pre *object.ListNode = nil
	curr, next := a, a
	for curr != b {
		next = curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}
