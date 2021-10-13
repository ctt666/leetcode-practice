package dynamic

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	f1, f2 := 1, 2
	for i := 3; i < n; i++ {
		if i%2 == 1 {
			f1 = f1 + f2
		} else {
			f2 = f1 + f2
		}
	}
	return f1 + f2
}
