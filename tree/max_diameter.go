package tree

import (
	"algo_study/object"
	"algo_study/util"
)

var maxDiameter int = 0

func DiameterOfBinaryTree(root *object.TreeNode) int {
	maxDepth(root)
	return maxDiameter
}

func maxDepth(root *object.TreeNode) int {
	if root == nil {
		return 0
	}
	leftMaxDepth := maxDepth(root.Left)
	rightMaxDepth := maxDepth(root.Right)
	maxDiameterTmp := leftMaxDepth + rightMaxDepth
	maxDiameter = util.Max(maxDiameter, maxDiameterTmp)
	return util.Max(leftMaxDepth, rightMaxDepth) + 1
}
