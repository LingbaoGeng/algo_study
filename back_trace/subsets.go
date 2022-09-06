package back_trace

var ans2 [][]int
var track []int

func Subsets(nums []int) [][]int {
	ans2 = make([][]int, 0)
	track = make([]int, 0)
	backTrack2(nums, 0)
	return ans2
}

func backTrack2(nums []int, start int) {
	temp := make([]int, len(track))
	copy(temp, track)
	ans2 = append(ans2, temp)
	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])
		backTrack2(nums, i+1)
		track = track[:len(track)-1]
	}
}
