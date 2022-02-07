package string

import "strings"

func lengthOfLastWord(s string) int {
	arrays := strings.Split(s, " ")
	for i := len(arrays) - 1; i >= 0; i-- {
		if strings.Trim(arrays[i], " ") == "" {
			continue
		}
		return len(arrays[i])
	}
	return 0
}
