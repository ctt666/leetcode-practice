package divide

import (
	"fmt"
	"testing"
)

func solveNQueens(n int) (result [][]string) {
	if n < 1 {
		return nil
	}
	col := map[int]bool{}
	pie := map[int]bool{}
	na := map[int]bool{}
	queen := make([]int, n)

	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			result = append(result, generate(queen))
		}
		for j := 0; j < n; j++ {
			if col[j] || pie[row-j] || na[row+j] {
				continue
			}
			col[j] = true
			pie[row-j] = true
			na[row+j] = true
			queen[row] = j
			dfs(row + 1)
			col[j] = false
			pie[row-j] = false
			na[row+j] = false
			//queen[row] = -1
		}
	}
	dfs(0)
	return
}

func generate(queen []int) (result []string) {
	l := len(queen)
	for _, v := range queen {
		strByte := make([]byte, l)
		for i := 0; i < l; i++ {
			if i == v {
				strByte[i] = 'Q'
			} else {
				strByte[i] = '.'
			}
		}
		result = append(result, string(strByte))
	}
	return
}

func TestI(t *testing.T) {
	result := solveNQueens(7)
	for _, v := range result {
		fmt.Println(v)
	}
}
