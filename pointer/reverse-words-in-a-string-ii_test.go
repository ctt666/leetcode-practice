package pointer

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		words[i] = reverse([]byte(word))
	}
	return strings.Join(words, " ")
}

func reverse(s []byte) string {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
	return string(s)
}
