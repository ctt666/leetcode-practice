package main

//编码判断，判断编码是否有效，有效输出有效位组成的编码十进制
/**
E980A5：
1110 1001  E9--前三位表示有几个编码，其他为有效位；只有一个编码则首位为0
1000 0000  80--前两位为10，其他为有效位
1010 0101  A5
*/

import (
	"fmt"
)

func main1() {
	var s string
	fmt.Scan(&s)
	var raw []int = make([]int, len(s)/2)

	//得到编码
	for i, _ := range s {
		addition := 1
		if (i+1)%2 == 1 {
			addition = 16
		}
		if s[i] >= '0' && s[i] <= '9' {
			raw[i/2] += int(s[i]-'0') * addition
		}
		if s[i] >= 'A' && s[i] <= 'F' {
			raw[i/2] += (int(s[i]-'A') + 10) * addition
		}
	}
	n := len(raw)
	k := 1 << 7
	//只有一个字节标志位校验
	if n == 1 {
		if raw[0]&k != 0 {
			fmt.Println(-1)
		} else {
			fmt.Println(raw[0])
		}
		return
	}
	result := 0
	//校验第一个字节的标志位
	for i := 0; i < n; i++ {
		if raw[0]&k == 0 {
			fmt.Println(-1)
			return
		}
		k = k >> 1
	}
	//逐个字节拼接计算有效位--十进制
	result += (255 >> n) & raw[0]
	k = 1 << 7
	for i := 1; i < n; i++ {
		fmt.Println(raw[i] & k)
		//后面字节标志位校验
		if raw[i]&k != 128 {
			fmt.Println(-1)
			return
		}
		//拼接计算有效位
		result *= 64
		result += (255 >> 2) & raw[i]
	}
	fmt.Println(result)
}
