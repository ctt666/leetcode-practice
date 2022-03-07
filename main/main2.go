package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main2() {
	as := make([]int, 3)
	bs := make([]int, 3)
	fmt.Scanf("%d,%d,%d", &as[0], &as[1], &as[2])
	fmt.Scanf("%d,%d,%d", &bs[0], &bs[1], &bs[2])
	bs_bak := make([]int, 3)
	//bs用于布阵
	copy(bs_bak, bs)
	//将两数组排序
	sort.Ints(as)
	sort.Ints(bs_bak)
	//empty判断：min(bs)>=max(as)
	if as[2] <= bs_bak[0] {
		fmt.Println("empty")
		return
	}
	//记录bs数据对应的索引
	mb := make(map[int]int, 3)
	for i := 0; i < 3; i++ {
		mb[bs[i]] = i
	}
	//记录ma数据状态
	ma := make(map[int]bool, 3)
	//记录as对战数据
	result := make([]int, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !ma[as[j]] && as[j] > bs[i] {
				ma[as[j]] = true
				result[i] = as[j]
				break
			}
		}
	}
	//存储返回结果
	var ans string
	//写入不能获胜的局
	for i := 0; i < 3; i++ {
		if result[i] == 0 {
			for j := 0; j < 3; j++ {
				if !ma[as[j]] {
					result[i] = as[j]
					break
				}
			}
		}
		ans += strconv.Itoa(result[i])
		if i < 2 {
			ans += ","
		}
	}

	fmt.Println(ans)
}

//func main() {
//	a1, a2, a3 := 0, 0, 0
//	b1, b2, b3 := 0, 0, 0
//
//	fmt.Scanf("%d,%d,%d", &a1, &a2, &a3)
//	fmt.Scanf("%d,%d,%d", &b1, &b2, &b3)
//
//}
