package bfs___dfs

func maxAreaOfIsland(grid [][]int) int {
	column := len(grid[0])
	row := len(grid)
	maxArea := 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			maxArea = max(maxArea, area(i, j, grid))
		}
	}
	return maxArea
}

func area(x, y int, grid [][]int) int {
	if grid[x][y] == 0 {
		return 0
	}
	row := len(grid)
	column := len(grid[0])
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	sum := 1
	grid[x][y] = 0
	for i := 0; i < 4; i++ {
		x1 := x + dx[i]
		y1 := y + dy[i]
		if x1 >= 0 && x1 < row && y1 >= 0 && y1 < column {
			sum += area(x1, y1, grid)
		}
	}
	return sum
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
