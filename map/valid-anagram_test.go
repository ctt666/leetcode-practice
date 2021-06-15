package _map

import (
	"fmt"
	"testing"
)

func isAnagram(s string, t string) bool {
	//method1: sort and compare
	//method2: map
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]int)
	//删除条件
	for i := range s {
		m[s[i]]++
		m[t[i]]--
		if m[t[i]] == 0 {
			delete(m, t[i])
		}
		if m[s[i]] == 0 {
			delete(m, s[i])
		}
	}
	return len(m) == 0
}

func TestIsAnagram(t *testing.T) {
	s1 := "anagram"
	s2 := "nagaram"
	fmt.Println(isAnagram(s1, s2))
}
