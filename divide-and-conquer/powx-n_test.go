package divide

import (
	"fmt"
	"testing"
)

//recursion
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n%2 > 0 {
		return x * myPow(x, n-1)
	}
	return myPow(x*x, n/2)
}

//iteration
func myPowIteration(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var result float64 = 1
	for n > 0 {
		//result计算条件
		if n&1 > 0 {
			result *= x
		}
		//x n 无条件迭代计算
		x *= x
		n = n >> 1
	}
	return result
}

func TestPow(t *testing.T) {
	fmt.Println(myPowIteration(2, 10))
}
