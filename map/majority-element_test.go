package _map

import (
	"fmt"
	"testing"
)

func majorityElement(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
		if m[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}

func Test1(t *testing.T) {
	i := 1
	p := &i
	*p = 10
	fmt.Println(i)
}
