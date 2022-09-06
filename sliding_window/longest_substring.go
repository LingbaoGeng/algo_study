package sliding_window

import "algo_study/util"

func LengthOfLongestSubstring(s string) int {
	ans := 0
	window = make(map[byte]int)
	left, right := 0, 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++

		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		ans = util.Max(ans, right-left)
	}
	return ans
}
