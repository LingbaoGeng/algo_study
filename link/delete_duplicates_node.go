package link

import "algo_study/object"

// DeleteDuplicateNodes 删除有序链表中的重复结点
// 双指针
func DeleteDuplicateNodes(head *object.ListNode) *object.ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	//别忘了最后断开后续的重复结点
	slow.Next = nil
	return head
}
