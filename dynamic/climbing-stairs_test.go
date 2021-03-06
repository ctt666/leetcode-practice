package dynamic

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	f1, f2 := 1, 2
	for i := 3; i < n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f1 + f2
}
