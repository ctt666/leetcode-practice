package string

import (
	"fmt"
	"strings"
	"testing"
)

func reverseWords(s string) string {
	result := []string{}
	left, right, k := 0, len(s)-1, 0
	for left <= right {
		if s[left] == ' ' {
			left++
			continue
		}
		if s[right] == ' ' {
			right--
			continue
		}
		for k = right; k >= 0; k-- {
			if s[k] == ' ' {
				break
			}
		}
		result = append(result, s[k+1:right+1])
		right = k
	}
	return strings.Join(result, " ")
}

func TestReverseWords(t *testing.T) {
	s := "hello  world"
	fmt.Println(reverseWords(s))
	//  world hello
}
