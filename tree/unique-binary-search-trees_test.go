package tree

/**
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？
返回满足题意的二叉搜索树的种数。
*/
func numTrees(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}

//reversion
// func numTrees(n int) int {
//     if n <= 1 {
//         return 1
//     }
//     result := 0
//     for i := 1; i <= n; i++ {
//         result += numTrees(i - 1) * numTrees(n - i)
//     }
//     return result
// }
