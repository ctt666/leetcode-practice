package string

func generateParenthesis(n int) []string {
	result := []string{}
	if n == 0 {
		return result
	}
	var backTrack func(left, right int, combination []byte)
	backTrack = func(left, right int, combination []byte) {
		if right == n {
			result = append(result, string(combination))
			return
		}
		if left < n {
			backTrack(left+1, right, append(combination, '('))
		}
		if right < left {
			backTrack(left, right+1, append(combination, ')'))
		}
	}
	backTrack(0, 0, []byte{})
	return result
}
