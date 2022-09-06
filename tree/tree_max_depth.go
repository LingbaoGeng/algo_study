package tree

import (
	"algo_study/object"
	"algo_study/util"
)

var res, depth int = 0, 0

func MaxDepth(root *object.TreeNode) int {
	travel(root)
	return res
}

func travel(root *object.TreeNode) {
	if root == nil {
		return
	}
	depth++
	if root.Left == nil && root.Right == nil {
		res = util.Max(res, depth)
	}
	travel(root.Left)
	travel(root.Right)
	depth--
}
