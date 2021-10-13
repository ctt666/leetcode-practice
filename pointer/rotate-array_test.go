package pointer

import (
	"fmt"
	"testing"
)

func rotate(nums []int, k int) {
	n := len(nums)
	m := k % n
	if m == 0 {
		return
	}

	i, next, tmp := 0, 0, nums[0]
	for j := 0; j < n; j++ {
		next = (next + m) % n
		nums[next], tmp = tmp, nums[next]
		if next == i {
			i++
			next = i
			tmp = nums[i]
		}
	}
}

func TestArray(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(array, k)
	fmt.Println(array)
}
