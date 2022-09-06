package link

import "algo_study/object"

// MergeTwoLists 合并2个有序表
// 设置虚拟头结点，可以避免处理空指针的情况，降低代码的复杂性
func MergeTwoLists(list1 *object.ListNode, list2 *object.ListNode) *object.ListNode {
	dump := &object.ListNode{
		Val:  -1,
		Next: nil,
	}
	p := dump
	node1 := list1
	node2 := list2
	for node1 != nil && node2 != nil {
		if node1.Val <= node2.Val {
			p.Next = node1
			node1 = node1.Next
		} else {
			p.Next = node2
			node2 = node2.Next
		}
		p = p.Next
	}
	if node1 != nil {
		p.Next = node1
	}
	if node2 != nil {
		p.Next = node2
	}
	return dump.Next
}
