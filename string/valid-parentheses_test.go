package string

func isValid(s string) bool {
	m := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []byte{}
	for i, _ := range s {
		v, ok := m[s[i]]
		if !ok {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 || v != stack[len(stack)-1] {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
