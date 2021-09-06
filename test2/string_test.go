package test2

import (
	"fmt"
	"testing"
)

func stringToInt(input string) (result int) {
	base := 1
	for _, v := range input {
		p := 10 * base
		q := int(v - '0')
		result += p * q
		base *= 10
	}
	return result
}

func TestString(t *testing.T) {
	fmt.Println("12345")
}
