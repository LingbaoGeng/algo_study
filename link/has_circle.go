package link

import "algo_study/object"

// HasCircle 链表是否包含环
// 快慢指针
func HasCircle(head *object.ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// HasCirclePoint 寻找链表中环的起点
// 快慢指针找到相遇点，然后慢指针回到起点，同步next，再次相遇就是链表的起点
func HasCirclePoint(head *object.ListNode) *object.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
