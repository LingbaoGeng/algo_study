package tree

import (
	"algo_study/object"
	"fmt"
)

func LevelTravel(root *object.TreeNode) {
	var treeNodeSlice []*object.TreeNode
	if root != nil {
		treeNodeSlice = append(treeNodeSlice, root)
	}
	for len(treeNodeSlice) != 0 {
		size := len(treeNodeSlice)
		for i := 0; i < size; i++ {
			fmt.Printf("层序遍历节点: %d\n", treeNodeSlice[i].Val)
			if treeNodeSlice[i].Left != nil {
				treeNodeSlice = append(treeNodeSlice, treeNodeSlice[i].Left)
			}
			if treeNodeSlice[i].Right != nil {
				treeNodeSlice = append(treeNodeSlice, treeNodeSlice[i].Right)
			}
		}
		treeNodeSlice = treeNodeSlice[size:]
	}
}
