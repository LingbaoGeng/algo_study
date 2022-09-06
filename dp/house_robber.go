package dp

var memo []int
var memo1 []int

func Rob(nums []int) int {
	memo = make([]int, len(nums))
	for i, _ := range memo {
		memo[i] = -1
	}
	memo1 = make([]int, len(nums)+2)
	return dp4(nums)
}

func dp1(nums []int, start int) int {
	if start >= len(nums) {
		return 0
	}
	res := max(dp1(nums, start+1), nums[start]+dp1(nums, start+2))
	return res
}

func dp2(nums []int, start int) int {
	if start >= len(nums) {
		return 0
	}
	if memo[start] != -1 {
		return memo[start]
	}
	res := max(dp2(nums, start+1), nums[start]+dp2(nums, start+2))
	memo[start] = res
	return res
}

func dp3(nums []int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		memo1[i] = max(memo1[i+1], nums[i]+memo1[i+2])
	}
	return memo1[0]
}

func dp4(nums []int) int {
	dp_i_1, dp_i_2 := 0, 0
	dp_i := 0
	for i := len(nums) - 1; i >= 0; i-- {
		dp_i = max(dp_i_1, nums[i]+dp_i_2)
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}
	return dp_i
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
