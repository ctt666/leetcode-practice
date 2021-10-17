package test2

import (
	"fmt"
	"testing"
)

func stringToInt(input string) (result int) {
	base := 1
	for i := len(input) - 1; i >= 0; i-- {
		result += int(input[i]-'0') * base
		base *= 10
	}
	return result
}

func TestString(t *testing.T) {
	fmt.Println(stringToInt("12345"))
}
