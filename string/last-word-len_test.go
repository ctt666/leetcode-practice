package string

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	Scanf(&s)
	raw := strings.Split(s, " ")
	last := raw[len(raw)-1]
	fmt.Println(len(last))
}

func Scanf(input *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	*input = string(data)
}
