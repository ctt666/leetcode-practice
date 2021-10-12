package dynamic

func minimumTotal(triangle [][]int) int {
	row := len(triangle)
	if row == 0 {
		return 0
	}
	d := make([]int, len(triangle[row-1]))
	for i := 0; i < len(d); i++ {
		d[i] = triangle[row-1][i]
	}

	for i := row - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			d[j] = min(d[j], d[j+1]) + triangle[i][j]
		}
	}

	return d[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
