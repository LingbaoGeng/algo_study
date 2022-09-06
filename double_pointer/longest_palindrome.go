package double_pointer

// LongestPalindrome 寻找字符串最长回文串
// 回文串长度分为偶数和奇数，如果为偶数，需要从中间两个向两边扩展
// 回文串统一成从中间两个向两边扩展，偶数情况左右相等
func LongestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func palindrome(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
