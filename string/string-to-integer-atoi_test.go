package string

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func myAtoi(s string) (ans int) {
	//去除空格
	s = strings.Trim(s, " ")
	//确定开始-结束索引
	var start, stop int
	flag, i := 1, 0
	for i < len(s) {
		if i == 0 && (s[i] == '-' || s[i] == '+') {
			if s[i] == '-' {
				flag = -1
			}
			start = 1
			i++
			continue
		}
		if s[i] < '0' || s[i] > '9' {
			break
		}
		i++
	}

	stop = i - 1
	if i == 0 || (i == 1 && start == 1) {
		return 0
	}

	//遍历对应字符串
	ans = flag * int(s[start]-'0')

	for i := start + 1; i <= stop; i++ {
		if flag == 1 && (math.MaxInt32-int(s[i]-'0'))/10 < ans {
			return math.MaxInt32
		}
		if flag == -1 && (math.MinInt32+int(s[i]-'0'))/10 > ans {
			return math.MinInt32
		}

		ans = ans*10 + int(s[i]-'0')*flag
	}
	return ans
}

func TestAtoi(t *testing.T) {
	s := "-42"
	fmt.Println(myAtoi(s))
}
