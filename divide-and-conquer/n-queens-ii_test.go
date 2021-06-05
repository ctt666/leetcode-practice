package divide

func totalNQueens(n int) (result int) {
	if n < 1 {
		return 0
	}
	col := make(map[int]bool, n)
	pie := make(map[int]bool, 2*n-1)
	na := make(map[int]bool, 2*n-1)
	var dfs func(row int)
	dfs = func(row int) {
		if row >= n {
			result++
			return
		}
		for j := 0; j < n; j++ {
			if col[j] || pie[row+j] || na[row-j] {
				continue
			}
			col[j] = true
			pie[j+row] = true
			na[row-j] = true
			dfs(row + 1)
			col[j] = false
			pie[j+row] = false
			na[row-j] = false
		}
	}
	dfs(0)
	return
}
