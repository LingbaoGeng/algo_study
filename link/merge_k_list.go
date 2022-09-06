package link

import "algo_study/object"

// MergeKLists 合并k条链表
// 转化成合并2条链表
func MergeKLists(lists []*object.ListNode) *object.ListNode {
	var ans *object.ListNode = nil
	for i := 0; i < len(lists); i++ {
		ans = MergeTwoLists(ans, lists[i])
	}
	return ans
}

func MergeKLists1(lists []*object.ListNode) *object.ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*object.ListNode, left, right int) *object.ListNode {
	if left == right {
		return lists[left]
	}
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	return MergeTwoLists(merge(lists, left, mid), merge(lists, mid+1, right))
}
