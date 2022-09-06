package link

import "algo_study/object"

// Partition 分隔链表，小于x的排到前面，其余的保持顺序不变
// 两条链表，一条存小于x的结点，一条存剩下的，然后将两条合成一条
// 注意原链接释放掉
func Partition(head *object.ListNode, x int) *object.ListNode {
	dump1 := &object.ListNode{
		Val:  -1,
		Next: nil,
	}
	dump2 := &object.ListNode{
		Val:  -1,
		Next: nil,
	}
	p1, p2 := dump1, dump2
	p := head
	for p != nil {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		temp := p.Next
		p.Next = nil
		p = temp
	}
	p1.Next = dump2.Next
	return dump1.Next
}
