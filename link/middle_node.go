package link

import "algo_study/object"

// MiddleNode 链表的中间结点
// 快慢指针

func MiddleNode(head *object.ListNode) *object.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
