package dynamic

import (
	"fmt"
	"testing"
)

//状态定义： dp[i]存储以arr[i]结尾的最长上升子序列
//状态转移： dp[i] = dp[j]--maxLen + arr[i]; arr[j] < arr[i]
//初始化： dp[0] = arr[0]
//结果： max(dp[i])

//最长上升子序列
func LIS(arr []int) []int {
	//dp存储以i结尾，最长上升子序列
	dp := make([][]int, len(arr))
	for i := 0; i < len(dp); i++ {
		dp[i] = []int{arr[i]}
	}
	result := []int{}

	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] < arr[i] && len(dp[j])+1 >= len(dp[i]) {
				dp[i] = append(dp[j], arr[i])
				if len(dp[i]) > len(result) {
					result = dp[i]
				}
			}
		}
	}
	return result
}

func TestLIS(t *testing.T) {
	input := []int{2, 1, 5, 3, 6, 4, 8, 9, 7}
	fmt.Println(LIS(input))
}

func TestSlice(t *testing.T) {
	input := []int{2, 1, 5}
	addition := []int{2, 4}
	input = append(input, addition...)
	fmt.Println(input)
}
