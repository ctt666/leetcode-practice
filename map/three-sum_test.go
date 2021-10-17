package _map

import (
	"fmt"
	"sort"
	"testing"
)

//先排序
//左右哨兵
func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := range nums {
		//去重
		//new different from old equal one
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		//左右哨兵
		p, q := i+1, len(nums)-1
		for p < q {
			sum := nums[i] + nums[p] + nums[q]
			if sum < 0 {
				p++
				continue
			}
			if sum > 0 {
				q--
				continue
			}
			result = append(result, []int{nums[i], nums[p], nums[q]})
			//去重
			for p < q && nums[p] == nums[p+1] {
				p++
			}
			//去重
			for p < q && nums[q] == nums[q-1] {
				q--
			}
			//下一次迭代：中间可能出现多次匹配情况
			p++
			q--
		}
	}
	return result
}

func TestThreeSum(t *testing.T) {
	//nums := []int{-1,0,1,2,-1,-4}
	//fmt.Println(threeSum(nums))
	var p []int
	p = append(p, 1)
	fmt.Println(p)
}
