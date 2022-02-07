package string

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	match := len(strs[0]) - 1
	for k := 1; k < len(strs); k++ {
		for i := 0; i <= match; i++ {
			if len(strs[k]) <= i || strs[k][i] != strs[0][i] {
				match = i - 1
				break
			}
		}
	}
	if match == -1 {
		return ""
	}
	return strs[0][:match+1]
}
