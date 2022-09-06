package bfs

import "algo_study/object"

func MinDepth(root *object.TreeNode) int {
	if root == nil {
		return 0
	}
	minDepth := 1
	treeNodeSlice := make([]*object.TreeNode, 0)
	treeNodeSlice = append(treeNodeSlice, root)
	for len(treeNodeSlice) != 0 {
		size := len(treeNodeSlice)
		for i := 0; i < size; i++ {
			if treeNodeSlice[i].Left == nil && treeNodeSlice[i].Right == nil {
				return minDepth
			}
			if treeNodeSlice[i].Left != nil {
				treeNodeSlice = append(treeNodeSlice, treeNodeSlice[i].Left)
			}
			if treeNodeSlice[i].Right != nil {
				treeNodeSlice = append(treeNodeSlice, treeNodeSlice[i].Right)
			}
		}
		treeNodeSlice = treeNodeSlice[size:]
		minDepth++
	}
	return minDepth
}
