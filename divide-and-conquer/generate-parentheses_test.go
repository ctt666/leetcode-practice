package divide

type Self struct {
	List []string
}

func generateParenthesis(n int) []string {
	s := &Self{}
	s.helper(0, 0, n, "")
	return s.List
}

func (s *Self) helper(left, right, n int, result string) {
	if right == n {
		s.List = append(s.List, result)
		return
	}
	if left < n {
		s.helper(left+1, right, n, result+"(")
	}
	if right < left {
		s.helper(left, right+1, n, result+")")
	}
}
