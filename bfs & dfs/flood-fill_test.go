package bfs___dfs

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
