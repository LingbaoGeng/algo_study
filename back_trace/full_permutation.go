package back_trace

var ans [][]int

func FullPermutaiton(nums []int) [][]int {
	track := make([]int, 0)
	used := make([]bool, len(nums))
	backTrace(nums, track, used)
	return ans
}

func backTrace(nums []int, track []int, used []bool) {
	if len(track) == len(nums) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		ans = append(ans, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		track = append(track, nums[i])
		used[i] = true
		backTrace(nums, track, used)
		track = track[:len(track)-1]
		used[i] = false
	}
}
