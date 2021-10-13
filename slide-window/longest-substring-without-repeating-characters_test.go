package slide_window

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	left, right := 0, 0
	m := map[byte]int{}
	for right < len(s) {
		index, ok := m[s[right]]
		//只考虑区间范围内是否又重复值
		if ok && index >= left {
			left = index + 1
		}

		maxLen = max(maxLen, right-left+1)
		m[s[right]] = right
		right++
	}
	return maxLen
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
