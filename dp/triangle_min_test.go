package dp

//状态定义：dp[i][j]从下层节点到（i，j）的最小距离
//状态转移：dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
//初始化： dp[len-1] = triangle[len-1]
//结果：dp[0][0]

func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	dp := make([][]int, length)
	dp[length-1] = triangle[length-1]
	for i := length - 2; i >= 0; i-- {
		dp[i] = make([]int, len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}
	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
