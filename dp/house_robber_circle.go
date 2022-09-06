package dp

// 打家劫舍，数组首尾不能同时存在，相当于带环
func RobCircle(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(dpCircle(nums, 0, n-2), dpCircle(nums, 1, n-1))
}

func dpCircle(nums []int, start, end int) int {
	dp_i_1, dp_i_2 := 0, 0
	dp_i := 0
	for i := end; i >= start; i-- {
		dp_i = max(dp_i_1, nums[i]+dp_i_2)
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}
	return dp_i
}
