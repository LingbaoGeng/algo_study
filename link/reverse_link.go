package link

import "algo_study/object"

// ReverseList 反转链表
// 遍历 注意保存前一个结点和后一个结点
func ReverseList(head *object.ListNode) *object.ListNode {
	var pre *object.ListNode = nil
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}

// ReverseList1 反转链表
// 利用递归反转链表
func ReverseList1(head *object.ListNode) *object.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := ReverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// ReverseN 反转链表前N个结点
// 保存断开的下一个结点
var nextNode *object.ListNode = nil

func ReverseN(head *object.ListNode, n int) *object.ListNode {
	if n == 1 {
		nextNode = head.Next
		return head
	}
	last := ReverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = nextNode
	return last
}

// ReverseBetween 反转链表中起点left和终点right的结点
// 递归，转换成反转链表的前N个结点
func ReverseBetween(head *object.ListNode, left int, right int) *object.ListNode {
	if left == 1 {
		return ReverseN(head, right)
	}
	head.Next = ReverseBetween(head.Next, left-1, right-1)
	return head
}
