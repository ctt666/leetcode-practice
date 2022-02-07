package string

import "strconv"

func addBinary(a string, b string) string {
	ans := ""
	n := max(len(a), len(b))
	carry := 0
	for i := 0; i < n; i++ {
		if i < len(a) {
			carry += int(a[len(a)-i-1] - '0')
		}
		if i < len(b) {
			carry += int(b[len(b)-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry = carry / 2
	}
	if carry == 1 {
		ans = "1" + ans
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
