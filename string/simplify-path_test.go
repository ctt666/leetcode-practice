package string

import "strings"

func simplifyPath(path string) string {
	stack := []string{}
	pathList := strings.Split(path, "/")
	for _, v := range pathList {
		if v == "" || v == "." {
			continue
		}
		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, v)
	}
	if len(stack) == 0 {
		return "/"
	}
	ans := ""
	for _, v := range stack {
		ans += "/" + v
	}
	return ans
}
