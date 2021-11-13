package bfs___dfs

//问题：以某点为中心，向四周扩散并替换某数字
//典型广度/深度优先
//trick：使用二维数组存储点的坐标
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, 1, -1}
	oldColor := image[sr][sc]
	image[sr][sc] = newColor
	queue := [][]int{}
	queue = append(queue, []int{sr, sc})
	n := len(image)
	m := len(image[0])

	for i := 0; i < len(queue); i++ {
		node := queue[i]
		for j := 0; j < 4; j++ {
			x := node[0] + dx[j]
			y := node[1] + dy[j]
			if x >= 0 && x < n && y >= 0 && y < m && image[x][y] == oldColor {
				image[x][y] = newColor
				queue = append(queue, []int{x, y})
			}
		}
	}
	return image
}
