package dp

import (
	"fmt"
	"testing"
)

//最长上升子序列
func LIS(arr []int) []int {
	dp := make([][]int, len(arr))
	for i := 1; i < len(dp); i++ {
		dp[i] = []int{}
	}
	dp[0] = []int{arr[0]}
	//dp[0] = append(dp[0], arr[0])
	result := []int{}

	for i := 1; i < len(arr); i++ {
		tmp := dp[i]
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] && len(dp[j])+1 >= len(dp[i]) {
				//if len(dp[i]) == 0 || dp[i][0] > arr[j] {
				dp[i] = append(tmp, dp[j]...)
				dp[i] = append(dp[i], arr[i])
				//dp[i] = append(dp[j], arr[i])
				if len(dp[i]) > len(result) {
					result = dp[i]
				}
			}
			//}
		}
	}
	return result
}

func TestLIS(t *testing.T) {
	input := []int{2, 1, 5, 3, 6, 4, 8, 9, 7}
	fmt.Println(LIS(input))
}
