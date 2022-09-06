package sliding_window

import "math"

func MinWindow(s string, t string) string {
	need = make(map[byte]int)
	window = make(map[byte]int)
	left, right, valid := 0, 0, 0
	start := 0
	subLen := math.MaxInt
	for _, val := range t {
		need[byte(val)]++
	}
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < subLen {
				start = left
				subLen = right - left
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if subLen == math.MaxInt {
		return ""
	}
	return s[start : start+subLen]
}
