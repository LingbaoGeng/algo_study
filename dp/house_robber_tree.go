package dp

import "algo_study/object"

func RobTree(root *object.TreeNode) int {
	ans := dpTree(root)
	return max(ans[0], ans[1])
}

func dpTree(root *object.TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left := dpTree(root.Left)
	right := dpTree(root.Right)

	rob := root.Val + left[0] + right[0]
	not_rob := max(left[0], left[1]) + max(right[0], right[1])
	return []int{not_rob, rob}
}
