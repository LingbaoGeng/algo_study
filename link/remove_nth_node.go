package link

import "algo_study/object"

// RemoveNthFromEnd 删除链表的倒数第N个结点
// 涉及链表的最好是设置个头指针，这样不需要处理一些边界情况
func RemoveNthFromEnd(head *object.ListNode, n int) *object.ListNode {
	dump := &object.ListNode{Val: -1, Next: nil}
	dump.Next = head
	p2 := findNthNode(dump, n+1)
	p2.Next = p2.Next.Next
	return dump.Next
}

func findNthNode(head *object.ListNode, n int) *object.ListNode {
	p1, p2 := head, head
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}
