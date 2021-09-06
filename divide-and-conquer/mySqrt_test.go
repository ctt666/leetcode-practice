package divide

func mySqrt(x int) int {
	r := x
	for r*r > x {
		//牛顿迭代
		r = (r + x/r) / 2
	}
	return r
}
