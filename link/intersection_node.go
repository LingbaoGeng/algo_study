package link

import "algo_study/object"

// GetIntersectionNode 获取相交链表的第一个相交点
// 使2条链表长度相同，当一条链表结束时，接着遍历另一条链表
func GetIntersectionNode(headA, headB *object.ListNode) *object.ListNode {
	p1, p2 := headA, headB

	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}

		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
