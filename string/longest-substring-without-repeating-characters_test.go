package string

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	left, right := 0, 0
	m := make(map[byte]int, 16)
	for right < len(s) {
		index, ok := m[s[right]]
		if ok && index >= left {
			left = index + 1
		}

		maxLen = max(maxLen, right-left+1)
		m[s[right]] = right
		right++
	}
	return maxLen
}
