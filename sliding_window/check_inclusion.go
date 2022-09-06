package sliding_window

func CheckInclusion(s1 string, s2 string) bool {
	need = make(map[byte]int)
	window = make(map[byte]int)
	left, right := 0, 0
	valid := 0
	for _, val := range s1 {
		need[byte(val)]++
	}

	for right < len(s2) {
		c := s2[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			d := s2[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}
