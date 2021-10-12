package recursion

func combine(n int, k int) [][]int {
	result := [][]int{}
	tmp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if len(tmp)+(n-cur+1) < k {
			return
		}
		if len(tmp) == k {
			array := make([]int, k)
			copy(array, tmp)
			result = append(result, array)
			return
		}
		tmp = append(tmp, cur)
		dfs(cur + 1)
		tmp = tmp[:len(tmp)-1]
		dfs(cur + 1)
	}
	dfs(1)
	return result
}
